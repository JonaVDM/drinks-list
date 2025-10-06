<script lang="ts">
  import { PB } from "./lib/pocketbase";
</script>

<main>
  {#await PB.collection("drinks").getFullList() then drinks}
    <pre>{JSON.stringify(drinks, null, 2)}</pre>
  {/await}
</main>
