```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of virtual machine image SKUs for the specified location, edge zone, publisher, and offer.
 *
 * @summary Gets a list of virtual machine image SKUs for the specified location, edge zone, publisher, and offer.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineImageExamples/VirtualMachineImagesEdgeZone_ListSkus_MaximumSet_Gen.json
 */
async function virtualMachineImagesEdgeZoneListSkusMaximumSetGen() {
  const subscriptionId = "{subscription-id}";
  const location = "aaaaaaaaaaaa";
  const edgeZone = "aaaaa";
  const publisherName = "aaaaaaaaaaaa";
  const offer = "aaaaaaaaaaaa";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineImagesEdgeZone.listSkus(
    location,
    edgeZone,
    publisherName,
    offer
  );
  console.log(result);
}

virtualMachineImagesEdgeZoneListSkusMaximumSetGen().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_19.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.
