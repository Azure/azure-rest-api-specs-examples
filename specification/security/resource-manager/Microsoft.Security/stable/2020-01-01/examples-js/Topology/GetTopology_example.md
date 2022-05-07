Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a specific topology component.
 *
 * @summary Gets a specific topology component.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/Topology/GetTopology_example.json
 */
async function getTopology() {
  const subscriptionId = "3eeab341-f466-499c-a8be-85427e154bad";
  const resourceGroupName = "myservers";
  const ascLocation = "centralus";
  const topologyResourceName = "vnets";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.topology.get(resourceGroupName, ascLocation, topologyResourceName);
  console.log(result);
}

getTopology().catch(console.error);
```
