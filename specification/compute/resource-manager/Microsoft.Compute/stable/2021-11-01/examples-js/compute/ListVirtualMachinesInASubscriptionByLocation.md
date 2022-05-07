Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_17.3.1/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function listsAllTheVirtualMachinesUnderTheSpecifiedSubscriptionForTheSpecifiedLocation() {
  const subscriptionId = "{subscriptionId}";
  const location = "eastus";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualMachines.listByLocation(location)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsAllTheVirtualMachinesUnderTheSpecifiedSubscriptionForTheSpecifiedLocation().catch(
  console.error
);
```
