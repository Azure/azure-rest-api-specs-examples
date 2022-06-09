```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all service endpoint policy definitions in a service end point policy.
 *
 * @summary Gets all service endpoint policy definitions in a service end point policy.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/ServiceEndpointPolicyDefinitionList.json
 */
async function listServiceEndpointDefinitionsInServiceEndPointPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceEndpointPolicyName = "testPolicy";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.serviceEndpointPolicyDefinitions.listByResourceGroup(
    resourceGroupName,
    serviceEndpointPolicyName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listServiceEndpointDefinitionsInServiceEndPointPolicy().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.
