<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, minimum-scale=1, user-scalable=no, minimal-ui">
    <meta name="apple-mobile-web-app-capable" content="yes">
    <meta name="apple-mobile-web-app-status-bar-style" content="black">
    <title>e家修</title>

    <link rel="stylesheet" href="/static/css/style.css">
  </head>
  <body>


    <div class="container pb-20">
      <div class="row">
        <div class="md-content">
          <div class="md-display-1">
		  {{.Title}}
          </div>
		  {{str2html .Content}}

        </div>
      </div>

    </div>

   <div class="share-bottom">
     <div class="share-list">
       <div class="item weixin">
         <img src="/static/images/public_icon_wechat.png" alt="">
         <p>微信好友</p>
       </div>
       <div class="item frendly">
         <img src="/static/images/public_icon_circleoffriends.png" alt="">
         <p>微信朋友圈</p>
       </div>
       <div class="item weibo">
         <img src="/static/images/public_icon_microblog.png" alt="">
         <p>新浪微博</p>
       </div>
     </div>
     <div class="share-list share-list2">
       <div class="item qq">
         <img src="/static/images/public_icon_qqfriends.png" alt="">
         <p>QQ好友</p>
       </div>
       <div class="item kongjian">
         <img src="/static/images/public_icon_qqqzone.png" alt="">
         <p>QQ空间</p>
       </div>
       <div class="item link">
         <img src="/static/images/public_icon_copylink.png" alt="">
         <p>复制链接</p>
       </div>
     </div>
     <div class="bb"></div>
     <h4 class="cancel">取消</h4>
   </div>
    <div class="shadow">

    </div>
    <div class="close"><img src="/static/images/public_icon_shut1.png"/></div>
    <script type="text/javascript" src="/static/js/jquery.min.js"></script>

    <script type="text/javascript">
      jQuery(document).ready(function($) {
        $('.shared-icon').click(function(event) {
          $('.share-bottom').addClass('bottom-0');
          $('.shadow').addClass('shodow-height');
          $('.close').addClass('close-show');
        });

        $('.cancel').click(function(event) {
          $('.share-bottom').removeClass('bottom-0');
          $('.shadow').removeClass('shodow-height');
          $('.close').removeClass('close-show');
        });

        $('.close').click(function(event) {
          $('.share-bottom').removeClass('bottom-0');
          $('.shadow').removeClass('shodow-height');
          $('.close').removeClass('close-show');
        });
      });
    </script>
  </body>
</html>
