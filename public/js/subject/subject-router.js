'use strict';

angular.module('warnabroda')
  .config(['$routeProvider', function ($routeProvider) {
    $routeProvider
      .when('/subjects', {
        templateUrl: 'views/subject/subjects.html',
        controller: 'SubjectController',
        resolve:{
          resolvedSubject: ['Subject', function (Subject) {
            return Subject.query();
          }]
        }
      })
    }]);
