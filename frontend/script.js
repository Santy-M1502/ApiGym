document.getElementById("loginForm").addEventListener("submit", async function (e) {
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
        console.log(data)
        // Guardar token en localStorage (si usás JWT, por ejemplo)
        localStorage.setItem("token", data.token);
        // Redirigir a la nueva página
        window.location.href = "exercises.html";
      } else {
        alert("Login incorrecto");
      }
    });