const API_BASE = "http://localhost:8080";

async function signup() {
  const username = document.getElementById("signup-username").value;
  const email = document.getElementById("signup-email").value;
  const password = document.getElementById("signup-password").value;

  const res = await fetch(`${API_BASE}/signup`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ username, email, password }),
  });

  const data = await res.json();

  if (res.ok) {
    alert("Signup successful! Please login.");
    window.location.href = "login.html";
  } else {
    document.getElementById("auth-message").innerText = data.message;
  }
}

async function login() {
  const email = document.getElementById("login-email").value;
  const password = document.getElementById("login-password").value;

  const res = await fetch(`${API_BASE}/login`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ email, password }),
  });

  const data = await res.json();

  if (res.ok) {
    localStorage.setItem("token", data.token);
    window.location.href = "dashboard.html";
  } else {
    document.getElementById("auth-message").innerText = data.message;
  }
}

function getToken() {
  return localStorage.getItem("token");
}

function authGuard() {
  if (!getToken()) {
    window.location.href = "index.html";
  }
}

async function loadTasks() {
  authGuard();

  const res = await fetch(`${API_BASE}/tasks`, {
    headers: {
      Authorization: `Bearer ${getToken()}`,
    },
  });

  const tasks = await res.json();
  const list = document.getElementById("task-list");
  list.innerHTML = "";

  tasks.forEach((task) => {
    const li = document.createElement("li");

    li.innerHTML = `
      <span id="task-title-${task.id}">
        ${task.title}
      </span>

      <button onclick="editTask('${task.id}')" title="Edit Task">✏️</button>
      <button onclick="deleteTask('${task.id}')" title="Delete Task">🗑️</button>
    `;

    list.appendChild(li);
  });
}

async function createTask() {
  const titleInput = document.getElementById("task-title");
  const title = titleInput.value.trim();

  if (!title) return;

  await fetch(`${API_BASE}/tasks`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${getToken()}`,
    },
    body: JSON.stringify({ title }),
  });

  titleInput.value = "";
  loadTasks();
}

function editTask(taskId) {
  const span = document.getElementById(`task-title-${taskId}`);
  const oldTitle = span.innerText;

  span.innerHTML = `
    <input
      type="text"
      id="edit-input-${taskId}"
      value="${oldTitle}"
    />
    <button onclick="updateTask('${taskId}')" title="Save Changes">💾</button>
    <button onclick="loadTasks()" title="Cancel">❌</button>
  `;
}

async function updateTask(taskId) {
  const newTitle = document.getElementById(`edit-input-${taskId}`).value.trim();

  if (!newTitle) return;

  await fetch(`${API_BASE}/tasks/${taskId}`, {
    method: "PATCH",
    headers: {
      "Content-Type": "application/json",
      Authorization: `Bearer ${getToken()}`,
    },
    body: JSON.stringify({ title: newTitle }),
  });

  loadTasks();
}

async function deleteTask(taskId) {
  await fetch(`${API_BASE}/tasks/${taskId}`, {
    method: "DELETE",
    headers: {
      Authorization: `Bearer ${getToken()}`,
    },
  });

  loadTasks();
}

function logout() {
  localStorage.removeItem("token");
  window.location.href = "login.html";
}
