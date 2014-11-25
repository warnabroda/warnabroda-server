'use strict';

angular.module('warnabroda')
  .config(['$routeProvider', function ($routeProvider) {
    $routeProvider
      .when('/warnings', {
        templateUrl: 'views/warning/warnings.html',
        controller: 'WarningController',
        resolve:{
          resolvedWarning: ['Warning', function (Warning) {
            return Warning.query();
          }]
        }
      })
    }]);
