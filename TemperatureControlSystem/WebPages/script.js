var temperature = document.getElementById("Temperature")
var roomList = document.getElementById("RoomName")

function getTemperatureStatus(){
    var roomName =roomList.value
    var url = window.location.href.replace(/\/$/, "") + "/Thermometar/"+roomName;
    $.ajax({
        method: "POST",
        dataType: "json",
        contentType: "application/json",
        url: url,
        data: "{}",
        success: function(data) {
            temperature.innerHTML = data.Temperature
        },
    });
}

function getAirCondtionerSetatus(){
    var roomName =roomList.value
    var url = window.location.href.replace(/\/$/, "") + "/Airconditioner/"+roomName;
    $.ajax({
        method: "POST",
        dataType: "json",
        contentType: "application/json",
        url: url,
        data: "{}",
        success: function(data) {
            temperature.innerHTML = data.Temperature
        },
    });
}

function setAirCondtionerStatus(){
    var roomName =roomList.value
    var url = window.location.href.replace(/\/$/, "") + "/Airconditioner/"+roomName;
    $.ajax({
        method: "POST",
        dataType: "json",
        contentType: "application/json",
        url: url,
        data: "{}",
        success: function(data) {
            temperature.innerHTML = data.Temperature
        },
    });
}

function getRooms(){
    var roomName = roomList.value
    var url = window.location.href.replace(/\/$/, "") + "/GetRooms";
    $.ajax({
        method: "POST",
        dataType: "json",
        contentType: "application/json",
        url: url,
        data: "{}",
        success: function(resp) {
            console.log(resp);
            
        },
        error: function(error) {
            console.log(error);
        }
    });
}

getRooms();