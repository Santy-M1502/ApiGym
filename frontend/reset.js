const token = new URLSearchParams(window.location.search).get("token");

document.getElementById("resetForm").addEventListener("submit", async (e) => {
  e.preventDefault();
  const nuevaContrasena = document.getElementById("nuevaContrasena").value;

  if (!token) {
    alert("Token de recuperación no encontrado.");
    return;
  }


  const res = await fetch("http://localhost:8080/api/resetear-password", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ token, password: nuevaContrasena }),
    }); 


  const data = await res.json();
  alert(data.ok ? "¡Contraseña cambiada!" : data.error);
});