package publichtml

// InvalidVerificationLink ...
var InvalidVerificationLink = `
<!DOCTYPE html>
<!-- This page is displayed when someone navigates to a verify email or reset password link
     but their security token is wrong. This can either mean the user has clicked on a
     stale link (i.e. re-click on a password reset link after resetting their password) or
     (rarely) this could be a sign of a malicious user trying to tamper with your app.
 -->
<html>
  <head>
  <title>Invalid Link</title>
  <style type='text/css'>
    .container {
      border-width: 0px;
      display: block;
      font: inherit;
      font-family: 'Helvetica Neue', Helvetica;
      font-size: 16px;
      height: 30px;
      line-height: 16px;
      margin: 45px 0px 0px 45px;
      padding: 0px 8px 0px 8px;
      position: relative;
      vertical-align: baseline;
    }

    h1, h2, h3, h4, h5 {
      color: #0067AB;
      display: block;
      font: inherit;
      font-family: 'Open Sans', 'Helvetica Neue', Helvetica;
      font-size: 30px;
      font-weight: 600;
      height: 30px;
      line-height: 30px;
      margin: 0 0 15px 0;
      padding: 0 0 0 0;
    }
  </style>
  </head>
  <script type="text/javascript">
    function getUrlParameter(name) {
      name = name.replace(/[\[]/, '\\[').replace(/[\]]/, '\\]');
      var regex = new RegExp('[\\?&]' + name + '=([^&#]*)');
      var results = regex.exec(location.search);
      return results === null ? '' : decodeURIComponent(results[1].replace(/\+/g, ' '));
    };  

    window.onload = addDataToForm;

    function addDataToForm() {
      var username = getUrlParameter("username");
      document.getElementById("usernameField").value = username;

      // var appId = getUrlParameter("appId");
      document.getElementById("resendForm").action = 'RESEND_VERIFICATION_URL'
    }

  </script>

  <body> 
    <div class="container">
      <h1>Invalid Verification Link</h1>
        <form id="resendForm" method="POST" action="/resend_verification_email">
          <input id="usernameField" class="form-control" name="username" type="hidden" value="">
          <button type="submit" class="btn btn-default">Resend Link</button>
        </form>
    </div> 
  </body>
</html>
`
