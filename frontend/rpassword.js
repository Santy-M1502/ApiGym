const modal = document.getElementById("changeComplete")
const reset = document.getElementById("resetForm")
reset.addEventListener("submit", async function (e) {
      e.preventDefault();

      const urlParams = new URLSearchParams(window.location.search);
      const token = urlParams.get("token");
      const password = document.getElementById("newPassword").value;

      const response = await fetch("http://localhost:8080/api/resetear-password", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ token, password })
      });

      const data = await response.json();
      if (data.ok) {
        alert("Contraseña actualizada correctamente.");
        window.location.href = "/index.html";
      } else {
        alert("Error: " + data.error);
      }
    });

const toggleBtn = document.getElementById('toggleMode');
    toggleBtn.addEventListener('click', () => {
      document.body.classList.toggle('dark');
      toggleBtn.textContent = document.body.classList.contains('dark') ? '☀️' : '🌙';
    });