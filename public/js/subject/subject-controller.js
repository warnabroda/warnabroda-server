'use strict';

angular.module('warnabroda')
  .controller('SubjectController', ['$scope', '$modal', 'resolvedSubject', 'Subject',
    function ($scope, $modal, resolvedSubject, Subject) {

      $scope.subjects = resolvedSubject;

      $scope.create = function () {
        $scope.clear();
        $scope.open();
      };

      $scope.update = function (id) {
        $scope.subject = Subject.get({id: id});
        $scope.open(id);
      };

      $scope.delete = function (id) {
        Subject.delete({id: id},
          function () {
            $scope.subjects = Subject.query();
          });
      };

      $scope.save = function (id) {
        if (id) {
          Subject.update({id: id}, $scope.subject,
            function () {
              $scope.subjects = Subject.query();
              $scope.clear();
            });
        } else {
          Subject.save($scope.subject,
            function () {
              $scope.subjects = Subject.query();
              $scope.clear();
            });
        }
      };

      $scope.clear = function () {
        $scope.subject = {
          
          "name": "",
          
          "lang_key": "",
          
          "id": ""
        };
      };

      $scope.open = function (id) {
        var subjectSave = $modal.open({
          templateUrl: 'subject-save.html',
          controller: 'SubjectSaveController',
          resolve: {
            subject: function () {
              return $scope.subject;
            }
          }
        });

        subjectSave.result.then(function (entity) {
          $scope.subject = entity;
          $scope.save(id);
        });
      };
    }])
  .controller('SubjectSaveController', ['$scope', '$modalInstance', 'subject',
    function ($scope, $modalInstance, subject) {
      $scope.subject = subject;

      

      $scope.ok = function () {
        $modalInstance.close($scope.subject);
      };

      $scope.cancel = function () {
        $modalInstance.dismiss('cancel');
      };
    }]);
