'use strict';

angular.module('warnabroda')
  .factory('Warning', ['$resource', function ($resource) {
    return $resource('warnabroda/warnings/:id', {}, {
      'query': { method: 'GET', isArray: true},
      'get': { method: 'GET'},
      'update': { method: 'PUT'}
    });
  }]);
