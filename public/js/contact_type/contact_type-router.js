'use strict';

angular.module('warnabroda')
  .config(['$routeProvider', function ($routeProvider) {
    $routeProvider
      .when('/contact_types', {
        templateUrl: 'views/contact_type/contact_types.html',
        controller: 'Contact_typeController',
        resolve:{
          resolvedContact_type: ['Contact_type', function (Contact_type) {
            return Contact_type.query();
          }]
        }
      })
    }]);
