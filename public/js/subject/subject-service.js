'use strict';

angular.module('warnabroda')
  .factory('Subject', ['$resource', function ($resource) {
    return $resource('warnabroda/subjects/:id', {}, {
      'query': { method: 'GET', isArray: true},
      'get': { method: 'GET'},
      'update': { method: 'PUT'}
    });
  }]);
