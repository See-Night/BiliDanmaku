# BiliDanmaku

BiliDanmaku是一个记录[bilibili直播](https://live.bilibili.com/)弹幕的项目.你可以在这里下载弹幕记录.如果这里没有你需要的某个主播的弹幕记录,你可以在Github上提交Issues,我回头有空的时候会把直播间加进去的.

BiliDanmaku is a project for record [bilibili live](https://live.bilibili.com) danmaku(弹幕). You can download danmaku log in here. If not have UPs that you need, you can submit issues on Github, and I will add his/her room at a later.

<ul id="danmaku_list">
</ul>



<script src="public/jquery.min.js" type="text/javascript"></script>
<script type="text/javascript">
    $().ready(() => {
        let danmaku_list = $('#danmaku_list');
        let api = 'https://api.github.com/repos/See-Night/BiliDanmaku/contents/?ref=logs';
        $.get(api, (res) => {
            for (let r in res) {
                danmaku_list.append(
                    $('<li></li>')
                    .text(res[r].name)
                    .click(() => {
                        window.open(res[r].html_url)
                    })
                )
            }
        })
    })
</script>





