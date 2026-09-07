import "./style.css";

// Interface untuk data Hobby dari GraphQL.
interface Hobby {
  id: string;
  name: string;
}

// Interface untuk data Person dari GraphQL.
interface Person {
  id: string;
  name: string;
  age: number;
  gender: string;
  job: string;
  phone: string;
  email: string;
  address: string;
  city: string;
  country: string;
  hobbies: Hobby[];
}

// Interface untuk struktur response GraphQL.
interface GraphQLResponse {
  data?: {
    people: Person[];
  };
  errors?: {
    message: string;
  }[];
}

// Menampilkan struktur awal halaman.
document.querySelector<HTMLDivElement>("#app")!.innerHTML = `
  <main class="container">
    <header class="header">
      <h1>Public Profile</h1>
      <p>Discover people and their hobbies</p>
    </header>

    <section class="controls">
      <div class="search-box">
        <input
          type="number"
          id="person-id"
          placeholder="Enter person ID"
          min="1"
        />

        <button id="search-button">
          Search Person
        </button>
      </div>

      <button id="random-button" class="random-button">
        Random Person
      </button>
    </section>

    <section class="profile-section">
      <h2>People</h2>

      <div id="people-list" class="people-list">
        <p>Loading people...</p>
      </div>
    </section>
  </main>
`;

// Mengambil elemen tempat daftar people akan ditampilkan.
const peopleList = document.querySelector<HTMLDivElement>("#people-list")!;

// Mengirim query GraphQL ke backend.
async function graphqlRequest<T>(query: string): Promise<T> {
  const response = await fetch("http://localhost:8080/query", {
    method: "POST",

    headers: {
      "Content-Type": "application/json",
    },

    body: JSON.stringify({
      query,
    }),
  });

  const result = await response.json();

  // Mengecek error dari GraphQL.
  if (result.errors) {
    throw new Error(result.errors[0].message);
  }

  return result.data;
}

// Mengubah data person menjadi HTML card.
function renderPerson(person: Person): string {
  return `
    <article class="person-card">
      <h3>${person.name}</h3>

      <p><strong>Age:</strong> ${person.age}</p>
      <p><strong>Gender:</strong> ${person.gender}</p>
      <p><strong>Job:</strong> ${person.job}</p>
      <p><strong>City:</strong> ${person.city}</p>

      <div class="hobbies">
        <strong>Hobbies:</strong>

        <div class="hobby-list">
          ${
            person.hobbies.length > 0
              ? person.hobbies
                  .map((hobby) => `<span class="hobby">${hobby.name}</span>`)
                  .join("")
              : "<span>No hobbies</span>"
          }
        </div>
      </div>
    </article>
  `;
}

// Mengambil satu person berdasarkan ID.
async function getPersonByID(id: string): Promise<void> {
  try {
    peopleList.innerHTML = "<p>Loading person...</p>";

    const data = await graphqlRequest<{
      person: Person;
    }>(`
      query {
        person(id: "${id}") {
          id
          name
          age
          gender
          job
          phone
          email
          address
          city
          country

          hobbies {
            id
            name
          }
        }
      }
    `);

    peopleList.innerHTML = renderPerson(data.person);
  } catch (error) {
    peopleList.innerHTML = `
      <p class="error">
        Failed to find person.
      </p>
    `;

    console.error(error);
  }
}

// Mengambil satu person secara random.
async function getRandomPerson(): Promise<void> {
  try {
    peopleList.innerHTML = "<p>Loading random person...</p>";

    const data = await graphqlRequest<{
      randomPerson: Person;
    }>(`
      query {
        randomPerson {
          id
          name
          age
          gender
          job
          phone
          email
          address
          city
          country

          hobbies {
            id
            name
          }
        }
      }
    `);

    peopleList.innerHTML = renderPerson(data.randomPerson);
  } catch (error) {
    peopleList.innerHTML = `
      <p class="error">
        Failed to get random person.
      </p>
    `;

    console.error(error);
  }
}

// Mengambil data people dari GraphQL API.
async function getPeople(): Promise<void> {
  try {
    // Mengirim request POST ke endpoint GraphQL backend.
    const response = await fetch("http://localhost:8080/query", {
      method: "POST",

      headers: {
        "Content-Type": "application/json",
      },

      // Query GraphQL yang dikirim ke backend.
      body: JSON.stringify({
        query: `
          query {
            people {
              id
              name
              age
              gender
              job
              phone
              email
              address
              city
              country
              hobbies {
                id
                name
              }
            }
          }
        `,
      }),
    });

    // Mengubah response menjadi JSON.
    const result: GraphQLResponse = await response.json();

    // Mengecek apakah GraphQL mengembalikan error.
    if (result.errors) {
      throw new Error(result.errors[0].message);
    }

    // Mengambil data people dari response.
    const people = result.data?.people ?? [];

    // Jika tidak ada data.
    if (people.length === 0) {
      peopleList.innerHTML = "<p>No people found.</p>";
      return;
    }

    // Menampilkan setiap person sebagai card.
    peopleList.innerHTML = people
      .map((person) => renderPerson(person))
      .join("");
  } catch (error) {
    // Menampilkan error jika request gagal.
    peopleList.innerHTML = `
      <p class="error">
        Failed to load people.
      </p>
    `;

    console.error(error);
  }
}

// Mengambil elemen input dan tombol.
const personIDInput = document.querySelector<HTMLInputElement>("#person-id")!;
const searchButton =
  document.querySelector<HTMLButtonElement>("#search-button")!;
const randomButton =
  document.querySelector<HTMLButtonElement>("#random-button")!;

// Menjalankan pencarian berdasarkan ID.
searchButton.addEventListener("click", () => {
  const id = personIDInput.value.trim();

  if (!id) {
    alert("Please enter a person ID.");
    return;
  }

  getPersonByID(id);
});

// Mengambil person secara random.
randomButton.addEventListener("click", () => {
  getRandomPerson();
});

// Menampilkan seluruh people ketika halaman pertama dibuka atau Menjalankan fungsi untuk mengambil data ketika halaman dibuka.
getPeople();
