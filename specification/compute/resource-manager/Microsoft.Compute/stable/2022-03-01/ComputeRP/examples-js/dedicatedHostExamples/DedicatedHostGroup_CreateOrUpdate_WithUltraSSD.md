Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a dedicated host group. For details of Dedicated Host and Dedicated Host Groups please see [Dedicated Host Documentation] (https://go.microsoft.com/fwlink/?linkid=2082596)
 *
 * @summary Create or update a dedicated host group. For details of Dedicated Host and Dedicated Host Groups please see [Dedicated Host Documentation] (https://go.microsoft.com/fwlink/?linkid=2082596)
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/dedicatedHostExamples/DedicatedHostGroup_CreateOrUpdate_WithUltraSSD.json
 */
async function createOrUpdateADedicatedHostGroupWithUltraSsdSupport() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const hostGroupName = "myDedicatedHostGroup";
  const parameters = {
    additionalCapabilities: { ultraSSDEnabled: true },
    location: "westus",
    platformFaultDomainCount: 3,
    supportAutomaticPlacement: true,
    tags: { department: "finance" },
    zones: ["1"],
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.dedicatedHostGroups.createOrUpdate(
    resourceGroupName,
    hostGroupName,
    parameters
  );
  console.log(result);
}

createOrUpdateADedicatedHostGroupWithUltraSsdSupport().catch(console.error);
```
