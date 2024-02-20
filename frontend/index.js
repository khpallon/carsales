console.log('Hello')

function postCars(){
    fetch("/api/cars", {
  method: "POST",
  body: JSON.stringify({
    title: "VOLKSWAGEN",
    price: 500,
    description: "TYPE-R"
  }),
  headers: {
    "Content-type": "application/json; charset=UTF-8"
  }
});

}