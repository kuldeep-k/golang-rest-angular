var app = angular.module('mainApp', ['ngRoute']);
app.
  config(['$routeProvider', function($routeProvider) {
  $routeProvider.
    when('/home', {
        templateUrl: 'templates/home.html',   
        controller: HomeController, 
        title : 'Dashboard'
    }).
    when('/attribute', {
        templateUrl: 'templates/attribute.html',   
        controller: AttributeController, 
        title : 'Attributes'
    }).
    when('/category', {
        templateUrl: 'templates/category.html',   
        controller: CategoryController, 
        title : 'Categories'
    }).
    otherwise({redirectTo: '/'});
}]);

/*app.run(['$rootScope', function($rootScope) {
    $rootScope.$on('$routeChangeSuccess', function (event, current, previous) {
        $rootScope.title = current.$$route.title;
    });
}]);
*/
var loginApp = angular.module('loginApp', ['ngRoute']);
loginApp.
  config(['$routeProvider', function($routeProvider) {
  $routeProvider.
    when('/login', {
        templateUrl: 'templates/login.html',   
        controller: LoginController, 
        title : 'Login'
    }).
    otherwise({redirectTo: '/login'});
}]);

