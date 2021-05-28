# BiliDanmaku

BiliDanmaku是一个记录[bilibili直播](https://live.bilibili.com/)弹幕的项目，你可以在这里下载弹幕记录。如果这里没有你需要的某个主播的弹幕记录，你可以在Github上提交Issues，我回头有空的时候会把直播间加进去的。

<img src="https://i0.hdslb.com/bfs/face/6f4808c37c98a183452f590fedbe80fc45201f60.jpg" style="zoom:25%;" />

### 主播列表

<div id="danmaku_list"></div>

<script src='public/jquery.min.js'></script>
<script type="text/javascript">
    $().ready(() => {
        let danmaku_list = $('#danmaku_list');
        let api = 'https://kaguramea.net/biliDanmaku';
        danmaku_list.css({
            display: 'flex',
            'flex-direction': 'row',
            'flex-wrap': 'wrap',
            width: '100%'
        });
        $.get(api, (res) => {
            for (let r in res.data) {
       			danmaku_list.append(
                    $('<img>')
                    	.css({
                        	'border-radius': '25px',
                        	height: '50px',
                       	 	width: '50px',
                        	margin: '5px'
                    	})
                    	.attr({
                            src: res.data[r].cover
                        })
                );         
            }
        });
    });
</script>




