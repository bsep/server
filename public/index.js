// ;

function urlPath (path) {
  return path.replace(/\\/g, '/');
}

angular
.module('App', [])
.controller('TableCtrl', function ($scope, $http, $rootScope) {
  $scope.filterName = function (row) {
    var text = row.Name.toLowerCase();
    var filter = ($scope.$root.filterName ? $scope.$root.filterName : '').toLowerCase();
    return filter.trim().length === 0 ? true : text.indexOf(filter) > -1;
  };

  $scope.filterDesc = function (row) {
    var text = row.Description.toLowerCase();
    var filter = ($scope.$root.filterDesc ? $scope.$root.filterDesc : '').toLowerCase();
    return filter.trim().length === 0 ? true : text.indexOf(filter) > -1;
  };

  $scope.delete = function (row) {
    $http.get('/admin/delete?filepath=' + row.Path).then($rootScope.refreshTable);
  };

  $scope.install = function (row) {
    window.external.AddSearchProvider(window.location.origin + '/plugins/' + urlPath(row.Path));
  };

  $scope.edit = function (row) {
    $rootScope.$broadcast('editRow', row);
  };

  $rootScope.refreshTable();
})
.controller('FilterCtrl', function ($scope) {
  $scope.hidden = true;
  $scope.clear = function ($event) {
    var model = angular.element($($event.currentTarget).prev()[0]).data('$ngModelController');
    model.$setViewValue('');
    model.$render();
    $scope.hidden = true;
  };
})
.controller('PackCtrl', function ($scope, $http, $rootScope) {
  $scope.pack = {};

  var packModal = $('#pack-modal');
  var packUrl = $('#pack-url');

  packModal.on('shown.bs.modal', function () { packUrl.focus(); });

  $scope.close = function () {
    packModal.modal('hide');
    $scope.pack.url = '';
    $scope.success = undefined;
  };

  $scope.submit = function () {
    // todo: error check
    $http.get('/admin/pack?url=' + $scope.pack.url)
      .then(function (res) {
        if (res.data.hasOwnProperty('success')) {
          $scope.success = res.data.success;
          $scope.message = res.data.message;
        }
        if ($scope.success == true) {
          $scope.pack.url = '';
        }
      })
      .then($rootScope.refreshTable);
  };
})
.controller('AddCtrl', function ($scope, $http, $rootScope) {
  $scope.add = {};

  $scope.save = function ($event) {
    var payload = new FormData();
    payload.append('filename', $scope.add.filename);
    payload.append('content', $scope.add.content);
    $http({
      method: 'POST',
      url: '/admin/add',
      headers: { 'Content-Type': undefined },
      cache: false,
      data: payload,
    })
    .then(function (res) { console.log('Add:', res); })
    .then($rootScope.refreshTable);
  };
})
.controller('EditCtrl', function ($scope, $http) {
  $scope.content = '';

  $scope.$on('editRow', function(e, row) {
    $scope.row = row;
    $http.get('/plugins/' + urlPath(row.Path) + '?' + Date.now())
      .then(function (res) { $scope.content = res.data; })
  })

  $scope.save = function ($event) {
    var payload = new FormData();
    payload.append('filepath', $scope.row.Path);
    payload.append('content', $scope.content);
    $http({
      method: 'POST',
      url: '/admin/edit',
      headers: { 'Content-Type': undefined },
      cache: false,
      data: payload,
    }).then(function (res) {
      console.log('Edit:', $scope.row.Path, '\nRes:', res);
      if (res.data.success) {
        console.log('Edited successfully:', $scope.row.Path);
      }
    })
  };
})
.run(function ($rootScope, $http) {
  $rootScope.table = [];
  $rootScope.refreshTable = function () {
    return $http.get('/list').then(function (res) { $rootScope.table = res.data; })
  };

  function toggleClass (e) { $(e.currentTarget).toggleClass('active'); }
  $('.navbar-nav li').hover(toggleClass, toggleClass);
})
