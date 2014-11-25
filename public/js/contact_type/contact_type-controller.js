'use strict';

angular.module('warnabroda')
  .controller('Contact_typeController', ['$scope', '$modal', 'resolvedContact_type', 'Contact_type',
    function ($scope, $modal, resolvedContact_type, Contact_type) {

      $scope.contact_types = resolvedContact_type;

      $scope.create = function () {
        $scope.clear();
        $scope.open();
      };

      $scope.update = function (id) {
        $scope.contact_type = Contact_type.get({id: id});
        $scope.open(id);
      };

      $scope.delete = function (id) {
        Contact_type.delete({id: id},
          function () {
            $scope.contact_types = Contact_type.query();
          });
      };

      $scope.save = function (id) {
        if (id) {
          Contact_type.update({id: id}, $scope.contact_type,
            function () {
              $scope.contact_types = Contact_type.query();
              $scope.clear();
            });
        } else {
          Contact_type.save($scope.contact_type,
            function () {
              $scope.contact_types = Contact_type.query();
              $scope.clear();
            });
        }
      };

      $scope.clear = function () {
        $scope.contact_type = {
          
          "name": "",
          
          "lang_key": "",
          
          "id": ""
        };
      };

      $scope.open = function (id) {
        var contact_typeSave = $modal.open({
          templateUrl: 'contact_type-save.html',
          controller: 'Contact_typeSaveController',
          resolve: {
            contact_type: function () {
              return $scope.contact_type;
            }
          }
        });

        contact_typeSave.result.then(function (entity) {
          $scope.contact_type = entity;
          $scope.save(id);
        });
      };
    }])
  .controller('Contact_typeSaveController', ['$scope', '$modalInstance', 'contact_type',
    function ($scope, $modalInstance, contact_type) {
      $scope.contact_type = contact_type;

      

      $scope.ok = function () {
        $modalInstance.close($scope.contact_type);
      };

      $scope.cancel = function () {
        $modalInstance.dismiss('cancel');
      };
    }]);
