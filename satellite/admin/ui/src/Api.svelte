<script lang="ts">
  import api from "./api";

  // types used for mapping the API operations to an auto generated UI
  type Operation = {
    name: string;
    operation: (...a: any) => any;
  };

  type GroupOperation = {
    name: string;
    operations: Operation[];
  };

  // initialize the available API Operations
  const apiOperations: GroupOperation[] = (() => {
    let gops = [];

    for (const g of Object.keys(api.operations)) {
      let gop = {
        name: g,
        operations: [],
      };

      for (const o of Object.keys(api.operations[g])) {
        gop.operations.push({
          name: o,
          operation: api.operations[g][o],
        });
      }

      gops.push(gop);
    }

    return gops;
  })();

  let token = "";
  let selectedGroupOp: Operation[];
</script>

<p>
  In order to work with the API you have to set the authentication token in the
  input box before executing any operation
</p>
<p>Token: <input bind:value={token} type="password" size="48" /></p>
<p>
  Operation:
  <select bind:value={selectedGroupOp}>
    <option selected />
    {#each apiOperations as group}
      <option value={group.operations}>{group.name}</option>
    {/each}
  </select>
  {#if selectedGroupOp}
    <select>
      <option selected />
      {#each selectedGroupOp as ops}
        <option>{ops.name}</option>
      {/each}
    </select>
  {/if}
</p>
