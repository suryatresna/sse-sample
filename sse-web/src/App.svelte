<script>
	export let name;
	import Card from './Card.svelte';
	import { backOut } from 'svelte/easing';
	import { crossfade } from 'svelte/transition';
	let dataNew = [];
	let dataOld = [];
	let indexes = {};
	function add(node, index) {
		indexes[index] = node;
		return {
			update(newIndex) {
				delete indexes[index];
				indexes[newIndex] = node;
			}
		}
	}
	const [send, receive] = crossfade({
		duration: d => Math.sqrt(d*600),
		fallback(node, params) {
			const style = getComputedStyle(node);
			const transform = style.transform === 'none' ? '' : style.transform;
			
			return {
				duration: 200,
				easing: backOut,
				css: t => `
					transform: ${transform} scale(${t});
					opacity: ${t};
					display: none;
				`
			};
		}
	})
	const sse = new EventSource('http://143.198.216.232:8080/sse');
	let uid = 1;
	sse.addEventListener("postupdate", e => {
		// console.log(e.data);
		let obj = JSON.parse(e.data);
		
		dataOld = [...dataNew, ...dataOld];
		dataNew = [obj];
		// console.log(obj)
		// control length data to prevent overheat browser
		if (dataOld.length > 10) {
			dataOld = dataOld.slice(0,9);
		}
	})
	sse.addEventListener("likeupdate", e => {
		// console.log(e.data);
		let obj = JSON.parse(e.data);
		for(var i in dataOld) {
			if (dataOld[i].id == obj.post_id) {
				dataOld[i].content.footer.like.value = obj.total_like;
				dataOld[i].content.footer.like.fmt = obj.total_like_human;
				// console.log(dataOld);
				break;
			}
		}
		for(var i in dataNew) {
			if (dataNew[i].id == obj.post_id) {
				dataNew[i].content.footer.like.value = obj.total_like;
				dataNew[i].content.footer.like.fmt = obj.total_like_human;
				// console.log(dataOld);
				break;
			}
		}
	})
	sse.addEventListener("commentupdate", e => {
		// console.log(e.data);
		let obj = JSON.parse(e.data);
		for(var i in dataOld) {
			if (dataOld[i].id == obj.post_id) {
				dataOld[i].content.footer.comment.value = obj.total_comment;
				dataOld[i].content.footer.comment.fmt = obj.total_comment_human;
				// console.log(dataOld);
				break;
			}
		}
		for(var i in dataNew) {
			if (dataNew[i].id == obj.post_id) {
				dataNew[i].content.footer.comment.value = obj.total_comment;
				dataNew[i].content.footer.comment.fmt = obj.total_comment_human;
				// console.log(dataOld);
				break;
			}
		}
	})
</script>

<main>
	<h1>Hello {name}!</h1>
	<p>Sample using Server-Sent Event, The page will automatically update by itself.</p>
	<section>
			{#each dataNew as item (item.id)}
				<div
					style="display:inline-box" 
					use:add={item}
					in:receive="{{key: item.id}}"
					>
					<Card 
						image={item.content.body.media[0].thumbnail}
						title={item.content.header.avatarTitle + " num: " + item.id}
						description={item.content.body.caption.text}
						likeCount={item.content.footer.like.fmt}
						commentCount={item.content.footer.comment.fmt}/>
				</div>
			{/each}
			{#each dataOld as item (item.id)}
				<div 
					use:add={dataOld}
					out:send="{{key: item.id}}"
					>
					<Card 
						image={item.content.body.media[0].thumbnail}
						title={item.content.header.avatarTitle + " num: " + item.id}
						description={item.content.body.caption.text}
						likeCount={item.content.footer.like.fmt}
						commentCount={item.content.footer.comment.fmt}/>
				</div>
			{/each}
	</section>
</main>

<style>
	* :global(.card-media-16x9) {
		background-image: url(https://via.placeholder.com/320x180.png?text=16x9);
	}
	* :global(.card-media-square) {
		background-image: url(https://via.placeholder.com/320x320.png?text=square);
	}
	main {
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
		text-align: center;
	}
	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}
	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>