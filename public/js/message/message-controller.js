'use strict';

angular.module('warnabroda')
  .controller('MessageController', ['$scope', '$modal', 'resolvedMessage', 'Message',
    function ($scope, $modal, resolvedMessage, Message) {

      $scope.messages = resolvedMessage;

      $scope.create = function () {
        $scope.clear();
        $scope.open();
      };

      $scope.update = function (id) {
        $scope.message = Message.get({id: id});
        $scope.open(id);
      };

      $scope.delete = function (id) {
        Message.delete({id: id},
          function () {
            $scope.messages = Message.query();
          });
      };

      $scope.save = function (id) {
        if (id) {
          Message.update({id: id}, $scope.message,
            function () {
              $scope.messages = Message.query();
              $scope.clear();
            });
        } else {
          Message.save($scope.message,
            function () {
              $scope.messages = Message.query();
              $scope.clear();
            });
        }
      };

      $scope.clear = function () {
        $scope.message = {
          
          "name": "",
          
          "lang_key": "",
          
          "id": ""
        };
      };

      $scope.open = function (id) {
        var messageSave = $modal.open({
          templateUrl: 'message-save.html',
          controller: 'MessageSaveController',
          resolve: {
            message: function () {
              return $scope.message;
            }
          }
        });

        messageSave.result.then(function (entity) {
          $scope.message = entity;
          $scope.save(id);
        });
      };
    }])
  .controller('MessageSaveController', ['$scope', '$modalInstance', 'message',
    function ($scope, $modalInstance, message) {
      $scope.message = message;

      

      $scope.ok = function () {
        $modalInstance.close($scope.message);
      };

      $scope.cancel = function () {
        $modalInstance.dismiss('cancel');
      };
    }]);
