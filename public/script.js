window.onload = function() {
}

var map;

function initMap() {
  var mapDiv = document.getElementById('map');


  var patroPos = {lat: 50.2015177, lng: 14.8392227};

  map = new google.maps.Map(mapDiv, {
    center: patroPos,
    zoom: 11
  });


  addPinToMap(1,50.1646768,14.8865066, "Restaurace Na Jitrách");
  addPinToMap(2,50.2015263,14.8370454, "Klub Patro");
  addPinToMap(3,50.146487,14.9056533, "Kersko");
  addPinToMap(4,50.1464659,14.8378023, "Skanzen v Přerově nad Labem");
  addPinToMap(5,50.1806812,14.7936072, "Byšičky");
  addPinToMap(6,50.2033323,14.8387577, "Barokní zámek a zahrada");
  addPinToMap(7,50.1979621,14.8417838, "Penzion Alfa");
}

function addPinToMap(letter, lat, lon, title) {

  var position = {lat: lat, lng: lon};
  var pinImage = new google.maps.MarkerImage("https://chart.googleapis.com/chart?chst=d_map_pin_letter_withshadow&chld="+letter+"|004375|FFFFFF");

  var marker = new google.maps.Marker({
    position: position,
    map: map,
    title: title,
    icon: pinImage
  });
}

