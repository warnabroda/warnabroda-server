'use strict';

angular.module('warnabroda')
  .factory('Contact_type', ['$resource', function ($resource) {
    return $resource('warnabroda/contact_types/:id', {}, {
      'query': { method: 'GET', isArray: true},
      'get': { method: 'GET'},
      'update': { method: 'PUT'}
    });
  }]);
