Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all Network Virtual Appliance Sites in a Network Virtual Appliance resource.
 *
 * @summary Lists all Network Virtual Appliance Sites in a Network Virtual Appliance resource.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/NetworkVirtualApplianceSiteList.json
 */
async function listAllNetworkVirtualApplianceSitesForAGivenNetworkVirtualAppliance() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const networkVirtualApplianceName = "nva";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.virtualApplianceSites.list(
    resourceGroupName,
    networkVirtualApplianceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllNetworkVirtualApplianceSitesForAGivenNetworkVirtualAppliance().catch(console.error);
```
