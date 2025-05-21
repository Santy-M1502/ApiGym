async function getProtectedData() {
      const token = localStorage.getItem("token");
      const response = await fetch("http://localhost:8080/ejercicios", {
        method: "GET",
        headers: {
          "Authorization": "Bearer " + token
        }
      });

      if (response.ok) {
        const data = await response.json();
        var lista = document.getElementById("data")
        lista.innerHTML = ""
        data.forEach((element, index) => {
          var div = document.createElement("div")
          div.classList.add("exercise_box")
          div.innerHTML = `
            <h3>Ejercicio: ${element.nombre}</h3>
            <p>Descripcion: ${element.descripcion}</p>
            <p>Repeticiones: ${element.repeticiones}</p>
          `
          const deleteButton = document.createElement('button');
          deleteButton.textContent = 'Eliminar';
          deleteButton.classList.add('delete-button')
          deleteButton.addEventListener('click', () => {
              console.log("Aca se elimina un elemento y se recarga la pagina")
              deleteExe(element.id)
              getProtectedData()
          })
          div.appendChild(deleteButton)
          lista.appendChild(div)
        ;
        });
      } else {
        alert("No autorizado. Redirigiendo al login.");
        window.location.href = "index.html";
      }
    }

getProtectedData();

const addExercise = document.getElementById("addForm")
addExercise.addEventListener("submit", (e)=>{
  e.preventDefault()
  var nombre = e.target[0].value
  var descripcion = e.target[1].value
  var repeticiones = e.target[2].value
  createExe(nombre, descripcion, repeticiones)
})

const createExe = (nombre, descripcion, repeticiones) =>{
  const token = localStorage.getItem("token");
  repeticiones = parseInt(repeticiones)
  fetch("http://localhost:8080/ejercicios",{
    method:"POST",
    headers: {
      "Authorization": "Bearer " + token,
      "Content-Type": "application/json"
    },
    body: JSON.stringify({nombre,descripcion,repeticiones})
  })
}

const deleteExe = (id) =>{
  const token = localStorage.getItem("token");
  fetch(`http://localhost:8080/ejercicios/${id}`,{
    method: "DELETE",
        headers: {
          "Authorization": "Bearer " + token,
          "Content-Type": "application/json"
        }
      })
}

document.addEventListener('DOMContentLoaded', () => {
  const toggleBtn = document.getElementById('toggleMode');
  const body = document.body;

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
