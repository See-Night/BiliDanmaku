$().ready(() => {
	let danmaku_list = $('#danmaku_list');
	let api = "https://kaguramea.net/biliDanmaku";
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
						onload: 'this.src = "https://kaguramea.net/media/danmaku/' + res.data[r].roomid + '.jpg"',
						alt: '' + res.data[r].name,
						src: 'https://kaguramea.net/media/danmaku/default.jpg'
					})
					.click(() => {
						window.open('https://github.com/See-Night/BiliDanmaku/tree/logs/' + res.data[r].roomid);
					})
			);
		}
	});

	let searchBtn = $($('#search').find('button'));
	let searchText = $($('#search').find('input'));
	searchBtn.click(() => {
		window.open('https://github.com/See-Night/BiliDanmaku/tree/logs/' + searchText.val());
	});
	searchText.keydown((e) => {
		if (e.keyCode == 13) {
			window.open('https://github.com/See-Night/BiliDanmaku/tree/logs/' + searchText.val());
		}
	});
});
