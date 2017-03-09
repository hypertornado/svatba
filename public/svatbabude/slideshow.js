window.onload = function() {
  initSlideshow();
}

var defaultSecondsToNextReload = 10;

function initSlideshow() {

  $(".slideshow_settings_interval").val(defaultSecondsToNextReload);

  $("#slideshow img").click( function() {
    if ($(".slideshow_settings").hasClass("slideshow_settings-visible")){
      $(".slideshow_settings").removeClass("slideshow_settings-visible");
    } else {
      $(".slideshow_settings").addClass("slideshow_settings-visible");
    }
  });

  $(".slideshow_settings_interval, .slideshow_settings_type").on("change", function (){
    nextReload = Date.now();
  });

  var nextReload = Date.now();

  setInterval(function(){
    if (nextReload < Date.now()) {
      loadData();
    }
  }, 100);

  loadData();

  function loadData() {

    var secondsToNextReload = parseInt(
      $(".slideshow_settings_interval").val()
    );

    nextReload = Date.now() + 1000 * secondsToNextReload;

    var data = {
      "type": $("#slideshow_settings_type").val()
    }

    $.ajax({
        url: "/api/slideshow",
        type: 'GET',
        data: data,
        cache: false,
        dataType: 'json',
        success: function(item) {
          $("#slideshow img").attr("src", item["url"]);
          $("#slideshow img").attr("src", item["url"]);
          $(".slideshow_text_content").text(item["description"]);
        },
        error: function() {
            console.error("error while loading files");
        }
    });
  }

}