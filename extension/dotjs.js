$.ajax({
  url: 'https://localhost:3131/'+location.hostname.replace(/^www\./,'')+'.js',
  dataType: 'text',
  success: function (data) {
    $(function() {
      eval(data)
    })
  },
  error: function(){
    console.log('No .js server found at localhost:3131 (with SSL)')
    $.ajax({
      url: 'http://localhost:3131/'+location.hostname.replace(/^www\./,'')+'.js',
      dataType: 'text',
      success: function (data) {
        $(function() {
          console.log('.js server found at localhost:3131 (without SSL)')
          eval(data)
        })
      },
      error: function() {
        console.log('No .js server found at localhost:3131 (without SSL)')
      }
    })
  }
})
