// === LOGIN Y OLVIDÉ CONTRASEÑA ===
document.getElementById("loginForm").addEventListener("submit", async function (e) {
  if (e.submitter.value == "send") {
    e.preventDefault();

    const email = document.getElementById("email").value;
    const password = document.getElementById("password").value;

    const response = await fetch("http://localhost:8080/login", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email, password })
    });

    console.log(JSON.stringify({ email, password }));

    if (response.ok) {
      const data = await response.json();
      localStorage.setItem("token", data.token);
      window.location.href = "exercises.html";
    } else {
      alert("Login incorrecto");
    }
  } else if (e.submitter.value == "nosend") {
    e.preventDefault();
    console.log("Olvida la contraseña");
  }
});


// === MODO OSCURO / CLARO ===
document.addEventListener("DOMContentLoaded", () => {
  const toggleBtn = document.getElementById("toggleMode");
  const body = document.body;

  if (localStorage.getItem("theme") === "dark") {
    body.classList.add("dark");
    toggleBtn.textContent = "☀️ Modo Claro";
  }

  toggleBtn.addEventListener("click", () => {
    body.classList.toggle("dark");
    const isDark = body.classList.contains("dark");
    toggleBtn.textContent = isDark ? "☀️ Modo Claro" : "🌙 Modo Oscuro";
    localStorage.setItem("theme", isDark ? "dark" : "light");
  });
});


// === MODAL: ABRIR Y CERRAR ===
document.getElementById("abrirModal").addEventListener("click", () => {
  document.getElementById("modalForgot").style.display = "flex";
});

document.querySelector(".modal .close").addEventListener("click", () => {
  document.getElementById("modalForgot").style.display = "none";
});

window.addEventListener("click", (e) => {
  const modal = document.getElementById("modalForgot");
  if (e.target === modal) {
    modal.style.display = "none";
  }
});


// === FORMULARIO: RECUPERAR CONTRASEÑA ===
document.getElementById("formForgot").addEventListener("submit", async (e) => {
  e.preventDefault();

  const email = document.getElementById("email-forgot").value.trim();

  if (!email.includes("@")) {
    alert("Por favor, ingresa un email válido");
    return;
  }

  console.log("Email capturado:", email);
  console.log("JSON enviado:", JSON.stringify({ email }));

  try {
    const res = await fetch("http://localhost:8080/verificar-email", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ email })
    });

    const data = await res.json();

    if (data.ok && data.existe) {
      alert("Correo enviado correctamente.");
      document.getElementById("modalForgot").style.display = "none";
    } else if (data.ok && !data.existe) {
      alert("Ese correo no está registrado.");
    } else {
      alert("Error: " + data.error);
    }
  } catch (err) {
    console.error("Error de red u otro:", err);
  }
});
