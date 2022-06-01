Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all available run commands for a subscription in a location.
 *
 * @summary Lists all available run commands for a subscription in a location.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/runCommandExamples/RunCommand_List.json
 */
async function virtualMachineRunCommandList() {
  const subscriptionId = "subid";
  const location = "SoutheastAsia";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachineRunCommands.list(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

virtualMachineRunCommandList().catch(console.error);
```
