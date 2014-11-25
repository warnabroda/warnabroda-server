'use strict';

angular.module('warnabroda')
  .config(['$routeProvider', function ($routeProvider) {
    $routeProvider
      .when('/messages', {
        templateUrl: 'views/message/messages.html',
        controller: 'MessageController',
        resolve:{
          resolvedMessage: ['Message', function (Message) {
            return Message.query();
          }]
        }
      })
    }]);
