require.config({
  paths: {
    'vs': 'assets/monaco-editor/vs'
  }
});
var editor = null;
require(['vs/editor/editor.main'], function () {
  editor = monaco.editor.create(document.getElementById('container'), {
    value: [
      'SELECT year, SUM(number) as num FROM [bigquery-public-data:usa_names.usa_1910_2013]',
      ' WHERE name = "William" GROUP BY year ORDER BY year '
    ].join('\n'),
    language: 'sql'
  });
});


var app = angular.module('poc-bigquery-geo', []);

app.controller('AppCtrl', ['$scope', '$http', function ($scope, $http) {
  this.data = null;
  this.exec = function (event) {
    this.showLoading = true;


    $http.get('/query?sql='+ editor.getValue()).then(function(body) {
      if (body && body.data && body.data.length > 0) {
        this.data = {
          rows: body.data,
          headers: Object.keys(body.data[0])
        }
      }
      this.showLoading = false;
    }.bind(this), function(err){
      this.showLoading = false;
    });

    // setTimeout(function () {
    //   this.data = {
    //     rows: [{
    //       name: 'ra',
    //       age: 29
    //     }],
    //     headers: ['name', 'age']
    //   }

    //   $scope.$apply();
    // }.bind(this), 2000);
    event.preventDefault();

  }

}]);
