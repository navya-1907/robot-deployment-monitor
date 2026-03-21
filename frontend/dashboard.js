async function loadData(){

    const res = await fetch("http://localhost:8080/logs")
    const data = await res.json()

    const table = document.getElementById("robotTable")

    table.innerHTML = `
    <tr>
        <th>Robot</th>
        <th>Battery</th>
        <th>Speed</th>
        <th>Lidar</th>
        <th>Position</th>
    </tr>`

    data.forEach(robot => {

        let batteryClass = "high"

        if(robot.battery < 30) batteryClass = "low"
        else if(robot.battery < 60) batteryClass = "medium"

        table.innerHTML += `
        <tr>
        <td>${robot.robot_id}</td>
        <td class="${batteryClass}">${robot.battery}%</td>
        <td>${robot.speed.toFixed(2)}</td>
        <td class="${robot.lidar_status === "NOISE" ? "noise" : "ok"}">
            ${robot.lidar_status}
        </td>
        <td>(${robot.x}, ${robot.y})</td>
        </tr>`
    })
}

window.onload = loadData;
setInterval(loadData, 2000);