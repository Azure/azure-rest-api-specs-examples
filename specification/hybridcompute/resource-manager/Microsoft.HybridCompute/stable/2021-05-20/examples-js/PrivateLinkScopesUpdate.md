```javascript
const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

async function privateLinkScopeUpdate() {
  const subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4919";
  const resourceGroupName = "my-resource-group";
  const scopeName = "my-privatelinkscope";
  const parameters = {
    location: "westus",
    tags: { tag1: "Value1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.privateLinkScopes.createOrUpdate(
    resourceGroupName,
    scopeName,
    parameters
  );
  console.log(result);
}

privateLinkScopeUpdate().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-hybridcompute_3.1.0-beta.1/sdk/hybridcompute/arm-hybridcompute/README.md) on how to add the SDK to your project and authenticate.
