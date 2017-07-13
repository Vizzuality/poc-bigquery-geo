require.config({
  paths: {
    'vs': 'assets/monaco-editor/vs'
  }
});
var editor = null;
require(['vs/editor/editor.main'], function () {
  editor = monaco.editor.create(document.getElementById('container'), {
    value: [
      'SELECT COUNT(*) FROM `alerts.fires`\n where viz_inside(longitude, latitude, \'[[[-55,49],[-39,49],[-39,53],[-55,53],[-55,49]]]\')'
    ].join('\n'),
    language: 'sql'
  });
});


var app = angular.module('poc-bigquery-geo', []);

app.controller('AppCtrl', ['$scope', '$http', function ($scope, $http) {
  this.data = null;
  this.count = 0;

  this.exampleIntersect = function($event) {
    $event.preventDefault();
    editor.setValue('SELECT COUNT(*) FROM `alerts.fires`\n where viz_intersect(the_geom, \'{"type":"Feature","properties":{},"geometry":{"type":"Polygon","coordinates":[[[-51.50390625,63.78248603116502],[-39.0234375,63.78248603116502],[-39.0234375,67.941650035336],[-51.50390625,67.941650035336],[-51.50390625,63.78248603116502]]]}}\')');
  }
  this.exampleInside = function($event) {
    $event.preventDefault();
    editor.setValue('SELECT COUNT(*) FROM `alerts.fires`\n where viz_inside(longitude, latitude, \'[[[-55,49],[-39,49],[-39,53],[-55,53],[-55,49]]]\')');
  }


  this.exec = function (event) {
    this.count = 0;
    this.showLoading = true;
    this.interval = setInterval(function() {
      this.count++;
      $scope.$apply();
    }.bind(this), 1000);

    $http.get('/query?sql='+ editor.getValue()).then(function(body) {
      if (body && body.data && body.data.length > 0) {
        this.data = {
          rows: body.data,
          headers: Object.keys(body.data[0])
        }
      }
      this.showLoading = false;
      clearInterval(this.interval);
    }.bind(this), function(err){
        alert('Query not valid \n', err.error);
      this.showLoading = false;
      clearInterval(this.interval);
    }.bind(this));
    event.preventDefault();

  }

}]);
