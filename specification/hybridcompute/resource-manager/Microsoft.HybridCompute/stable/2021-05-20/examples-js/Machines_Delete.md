```javascript
const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

async function deleteAMachine() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "myResourceGroup";
  const machineName = "myMachine";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.machines.delete(resourceGroupName, machineName);
  console.log(result);
}

deleteAMachine().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-hybridcompute_3.1.0-beta.1/sdk/hybridcompute/arm-hybridcompute/README.md) on how to add the SDK to your project and authenticate.
