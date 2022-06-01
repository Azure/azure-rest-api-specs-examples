Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all available virtual machine sizes that can be used to create a new virtual machine in an existing availability set.
 *
 * @summary Lists all available virtual machine sizes that can be used to create a new virtual machine in an existing availability set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/availabilitySetExamples/AvailabilitySets_ListAvailableSizes_MinimumSet_Gen.json
 */
async function availabilitySetsListAvailableSizesMinimumSetGen() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "rgcompute";
  const availabilitySetName = "aa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.availabilitySets.listAvailableSizes(
    resourceGroupName,
    availabilitySetName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

availabilitySetsListAvailableSizesMinimumSetGen().catch(console.error);
```
