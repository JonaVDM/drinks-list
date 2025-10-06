<script lang="ts">
  import { PB } from "./lib/pocketbase";
</script>

<main>
  <h1 class="text-2xl font-thin">Dingetje</h1>

  {#await PB.collection("drinks").getFullList() then drinks}
    <pre>{JSON.stringify(drinks, null, 2)}</pre>
  {/await}
</main>
