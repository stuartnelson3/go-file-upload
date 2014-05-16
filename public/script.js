$(function () {
  'use strict';

  var url = location.origin + "/upload";
  $('#fileupload').fileupload({
    url: url,
    dataType: "json",
    done: function (e, data) {
      $.each(data.result, function (index, fileName) {
        $('<p/>').text(fileName).appendTo('#files');
      });
    },
    progressall: function (e, data) {
      var progress = parseInt(data.loaded / data.total * 100, 10);
      $('#progress .progress-bar').css(
        'width',
        progress + '%'
      );
    }
  }).prop('disabled', !$.support.fileInput)
  .parent().addClass($.support.fileInput ? undefined : 'disabled');
});
