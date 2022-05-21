Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified ServiceEndpoint policy definitions.
 *
 * @summary Deletes the specified ServiceEndpoint policy definitions.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/ServiceEndpointPolicyDefinitionDelete.json
 */
async function deleteServiceEndpointPolicyDefinitionsFromServiceEndpointPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const serviceEndpointPolicyName = "testPolicy";
  const serviceEndpointPolicyDefinitionName = "testDefinition";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.serviceEndpointPolicyDefinitions.beginDeleteAndWait(
    resourceGroupName,
    serviceEndpointPolicyName,
    serviceEndpointPolicyDefinitionName
  );
  console.log(result);
}

deleteServiceEndpointPolicyDefinitionsFromServiceEndpointPolicy().catch(console.error);
```
