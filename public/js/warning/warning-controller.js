'use strict';

angular.module('warnabroda')
  .controller('WarningController', ['$scope', '$modal', 'resolvedWarning', 'Warning',
    function ($scope, $modal, resolvedWarning, Warning) {

      $scope.warnings = resolvedWarning;

      $scope.create = function () {
        $scope.clear();
        $scope.open();
      };

      $scope.update = function (id) {
        $scope.warning = Warning.get({id: id});
        $scope.open(id);
      };

      $scope.delete = function (id) {
        Warning.delete({id: id},
          function () {
            $scope.warnings = Warning.query();
          });
      };

      $scope.save = function (id) {
        if (id) {
          Warning.update({id: id}, $scope.warning,
            function () {
              $scope.warnings = Warning.query();
              $scope.clear();
            });
        } else {
          Warning.save($scope.warning,
            function () {
              $scope.warnings = Warning.query();
              $scope.clear();
            });
        }
      };

      $scope.clear = function () {
        $scope.warning = {
          
          "id_message": "",
          
          "id_contact_type": "",
          
          "contact": "",
          
          "sent": "",
          
          "message": "",
          
          "ip": "",
          
          "browser": "",
          
          "operating_system": "",
          
          "device": "",
          
          "raw": "",
          
          "created_by": "",
          
          "created_date": "",
          
          "last_modified_by": "",
          
          "last_modified_date": "",
          
          "lang_key": "",
          
          "id": ""
        };
      };

      $scope.open = function (id) {
        var warningSave = $modal.open({
          templateUrl: 'warning-save.html',
          controller: 'WarningSaveController',
          resolve: {
            warning: function () {
              return $scope.warning;
            }
          }
        });

        warningSave.result.then(function (entity) {
          $scope.warning = entity;
          $scope.save(id);
        });
      };
    }])
  .controller('WarningSaveController', ['$scope', '$modalInstance', 'warning',
    function ($scope, $modalInstance, warning) {
      $scope.warning = warning;

      
      $scope.created_dateDateOptions = {
        dateFormat: 'yy-mm-dd',
        
        
      };
      $scope.last_modified_dateDateOptions = {
        dateFormat: 'yy-mm-dd',
        
        
      };

      $scope.ok = function () {
        $modalInstance.close($scope.warning);
      };

      $scope.cancel = function () {
        $modalInstance.dismiss('cancel');
      };
    }]);
