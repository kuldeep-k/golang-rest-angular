function HomeController($scope) {

}
function AttributeController($scope, $http){
    /*$http.get('videoList.json').success(function(videoList) {
        $scope.data = videoList;
    });*/
    var url = "http://localhost:8085/attribute/list";                  
          
     $http.get(url).success( function(response) {
        attributes = [];
        var i = 1;
        for(key in response) {
            attributes.push({
                'sno' : i, 
                'name' : response[key].title, 
                'type' : response[key].attributetype, 
                'created' : response[key].createdDate, 
                'modified' : response[key].modifiedDate, 
                'status' : response[key].status, 
                'id' : response[key].id 
            });
            i++;
        }
        //$scope.breadcrumbTitle = "Attribute";
        //$scope.pageTitle = "Attribute";
        $scope.attributes = attributes;

     });
}
 
function CategoryController($scope,$http){
    var url = "http://localhost:8085/category/list";                      
          
     $http.get(url).success( function(response) {
        categories = [];
        var i = 1;
        for(key in response) {
            categories.push({
                'sno' : i, 
                'name' : response[key].name, 
                'created' : response[key].createdDate, 
                'modified' : response[key].modifiedDate, 
                'status' : response[key].status, 
                'id' : response[key].id 
            });
            i++;
        }
        //$scope.breadcrumbTitle = "Category";
        //$scope.pageTitle = "Category";
        $scope.categories = categories;

     });
}

function LoginController($scope, $http) {
    $scope.formData = {};

    $scope.loginRequest = function() {
        var url = "http://stryker-traquer.igenuine.net/mis/login";   
        $http({
              method  : 'POST',
              url     : url,
              data    : $.param($scope.formData),  // pass in data as strings
              headers : { 'Content-Type': 'application/x-www-form-urlencoded' }  // set the headers so angular passing info as form data (not request payload)
             })
              .success(function(data) {
                console.log(data);

                if (!data.success) {
                  // if not successful, bind errors to error variables
                  $scope.usernameError = data.message;
                } else {
                  // if successful, bind success message to message
                  $scope.message = data.message;
                }
              });
      };
}

