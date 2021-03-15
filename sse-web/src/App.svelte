<script>
	export let name;

	let messages =["Test","hello"];

	

</script>

<main>
	<h1>Hello {name}!</h1>
	<p>This is sample integration client SSE (Server Sent Event)</p>
	<ul>
	{#each messages as message, i}
		<li id={"list-"+i}>
			{message}
		</li>
	{/each}
	</ul>
</main>

<style>
	main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
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