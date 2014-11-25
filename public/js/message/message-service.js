'use strict';

angular.module('warnabroda')
  .factory('Message', ['$resource', function ($resource) {
    return $resource('warnabroda/messages/:id', {}, {
      'query': { method: 'GET', isArray: true},
      'get': { method: 'GET'},
      'update': { method: 'PUT'}
    });
  }]);
