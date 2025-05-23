document.getElementById("loginForm").addEventListener("submit", async function (e) {
  if (e.submitter.value == "send"){ 
    e.preventDefault();
      
      const email = document.getElementById("email").value;
      const password = document.getElementById("password").value;
      const response = await fetch("http://localhost:8080/login", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ "email": email, "password" : password })
      });
      console.log(JSON.stringify({email,password }))

      if (response.ok) {
        const data = await response.json();
        // Guardar token en localStorage (si usás JWT, por ejemplo)
        localStorage.setItem("token", data.token);
        // Redirigir a la nueva página
        window.location.href = "exercises.html";
      } else {
        alert("Login incorrecto");
      }
    }else if(e.submitter.value == "nosend"){
      e.preventDefault();
      console.log("Olvida la contrasena")
    }})
      

document.addEventListener('DOMContentLoaded', () => {
  const toggleBtn = document.getElementById('toggleMode');
  const body = document.body;

  // Aplica tema guardado
  if (localStorage.getItem('theme') === 'dark') {
    body.classList.add('dark');
    toggleBtn.textContent = '☀️ Modo Claro';
  }

  toggleBtn.addEventListener('click', () => {
    body.classList.toggle('dark');
    const isDark = body.classList.contains('dark');
    toggleBtn.textContent = isDark ? '☀️ Modo Claro' : '🌙 Modo Oscuro';
    localStorage.setItem('theme', isDark ? 'dark' : 'light');
  });
});

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
  }})

document.getElementById("formForgot").addEventListener("submit", (e) => {
  e.preventDefault();

  const emailInput = document.getElementById("email-forgot");
  if (emailInput.checkValidity()) {
    const email = emailInput.value;
    alert(`Se envió un correo a: ${email}`);
    document.getElementById("modalForgot").style.display = "none";
  } else {
    emailInput.reportValidity(); 
  }
});
