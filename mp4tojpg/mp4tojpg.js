

// forked from elk's "MP4ビデオをvideoからcanvasへ描画するテスト" http://jsdo.it/elk/gKdk
var theCanvas;
var context;
var videoElement;
var dataUrl;

function  drawVideo () {
      //Background
      context.fillStyle = '#ffffaa';
      context.fillRect(0, 0, theCanvas.width, theCanvas.height);
      //video
      context.drawImage(videoElement , 0, 0);
      window.requestAnimationFrame(drawVideo);
}
function capture() {

    dataUrl = theCanvas.toDataURL("image/jpeg");
    document.getElementById("svImg").setAttribute("href", dataUrl);
}

window.onload = function(){
    theCanvas = document.getElementById("canvas1");
    context = theCanvas.getContext("2d");
    videoElement = document.getElementById("video1");

     videoElement.play();
     drawVideo();

}
