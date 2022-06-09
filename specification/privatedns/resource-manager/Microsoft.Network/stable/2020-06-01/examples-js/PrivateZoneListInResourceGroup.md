```javascript
const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the Private DNS zones within a resource group.
 *
 * @summary Lists the Private DNS zones within a resource group.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2020-06-01/examples/PrivateZoneListInResourceGroup.json
 */
async function getPrivateDnsZoneByResourceGroup() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroup1";
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateZones.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getPrivateDnsZoneByResourceGroup().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-privatedns_3.0.1/sdk/privatedns/arm-privatedns/README.md) on how to add the SDK to your project and authenticate.
