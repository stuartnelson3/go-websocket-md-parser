var host = location.origin.replace(/^http/, 'ws') + '/markdown_preview';
var ws = new WebSocket(host);

$(document).on('keyup', 'textarea', function() {
  var body = $('textarea').val();
  ws.send(body);
});

ws.onmessage = function(e) {
  var container = document.querySelector('.preview');
  container.innerHTML = e.data;
  $('pre code').each(function(i, e) {hljs.highlightBlock(e)});
};
