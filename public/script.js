window.onload = function() {

}

function initMap() {
  var mapDiv = document.getElementById('map');

  //AIzaSyDgsEoKEfZJZJtwZe8_YtQ3CQFnOZ1xVuI


  var patroPos = {lat: 50.2015177, lng: 14.8392227};

  var map = new google.maps.Map(mapDiv, {
    center: patroPos,
    zoom: 11
  });

  var pinImage = new google.maps.MarkerImage("http://chart.apis.google.com/chart?chst=d_map_pin_letter&chld=%E2%80%A2%7C004375");

  var iconBase = 'https://maps.google.com/mapfiles/kml/shapes/';
  iconBase = "/logo/logo.png";
  var marker = new google.maps.Marker({
    position: patroPos,
    map: map,
    title: "Klub Patro",
    icon: pinImage
    //icon: iconBase// + 'schools_maps.png'
  });
}