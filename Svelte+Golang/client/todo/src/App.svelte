<script>
  import { onMount } from "svelte";
  import headers from "../header.js";

  let TodoList = [];
  let baseURL = "http://localhost:8080";
  let title = "";
  let action = [true, false];
  let completed = "";

  const toastManager = (message, toastColor) => {
    document.querySelector(".alert").innerHTML = message;
    document.querySelector(".alert").classList.add(toastColor);
  };

  async function handleDelete(id) {
    await fetch(`${baseURL}/delete/${id}`, {
      method: "PUT",
      headers: headers,
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
      })
      .catch((err) => {
        console.log(err);
      });
  }

  async function handleSubmit(e) {
    const body = { title, completed };
    console.log("body", JSON.stringify(body));
    const res = await fetch(`${baseURL}/create`, {
      method: "POST",
      headers: headers,
      body: JSON.stringify(body),
    })
      .then((res) => {
        console.log("ressdaffcaesZFCAAAAAAASD", res);
        res.json();
      })
      .catch((err) => {
        console.log(err);
      });
    console.log("rehiahiaiaihihas", res);
    toastManager("Task Added Successfully", "alert-success");
  }

  async function handleUpdate(id) {
    await fetch(`${baseURL}/update/${id}`, {
      method: "PUT",
      headers: headers,
    })
      .then((res) => res.json())
      .then((data) => {
        console.log("here is ", data);
      })
      .catch((err) => {
        console.log(err);
      });
  }

  onMount(async () => {
    await fetch(`${baseURL}/`)
      .then((res) => res.json())
      .then((data) => {
        console.log("data", data);
        TodoList = [...data, ...TodoList];
      })
      .catch((err) => {
        console.log(err);
      });
  });
</script>

<main>
  <div class="alert" role="alert" />
  <form on:submit|preventDefault={handleSubmit}>
    <h4>Enter the Task</h4>
    <input
      type="text"
      class="form-control"
      value={title}
      on:input={(e) => (title = e.target.value)}
    />
    <h4>Enter the Action</h4>
    {#each action as a}
      <label for="a"
        >{a ? "Completed" : "Not Completed"}
        <input type="radio" bind:group={completed} value={a} />
      </label>
    {/each}
    <button class="btn btn-success" type="submit">Button</button>
  </form>
  <table class="table table-striped table-hover">
    <thead>
      <tr>
        <th scope="col">#</th>
        <th scope="col">Name of Task</th>
        <th scope="col">Action</th>
        <th scope="col">Delete</th>
        <th scope="col">Update</th>
      </tr></thead
    >
    <tbody>
      {#each TodoList as todo}
        <tr>
          <th scope="row">{todo._id}</th>
          <td>{todo.title}</td>
          <td>{todo.completed}</td>
          <td
            ><button
              class="btn btn-danger"
              on:click={() => {
                handleDelete(todo._id);
              }}>Delete</button
            ></td
          >
          <td
            ><button
              class="btn btn-secondary"
              on:click={() => handleUpdate(todo._id)}>Update</button
            ></td
          >
        </tr>
      {/each}
    </tbody>
  </table>
</main>

<style>
  form {
    display: flex;
    flex-direction: column;
    justify-content: center;
    padding: 40px 10px 30px 10px;
  }
  input[type="text"] {
    margin: 10px;
  }
  button {
    margin: 10px 0 10px 0;
  }
</style>
