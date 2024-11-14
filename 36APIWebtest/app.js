fetch('http://localhost:5000/course')
    .then(function (response) {
        return response.json();
    })
    .then(function (data) {
        appendData(data);
    })
    .catch(function (err) {
        console.log('error: ' + err);
    });

function appendData(data) {
    var mainContainer = document.getElementById("myData");
    for (var i = 0; i < data.length; i++) {
        var div = document.createElement("div");
        div.innerHTML = 'CourseID: ' + data[i].id + ' ' + data[i].name + ' ' + data[i].price + ' ' + data[i].instructor;
        mainContainer.appendChild(div);
    }
}