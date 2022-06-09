```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about all public IP addresses on a cloud service level.
 *
 * @summary Gets information about all public IP addresses on a cloud service level.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/CloudServicePublicIpListAll.json
 */
async function listVmssPublicIP() {
  const subscriptionId = "subid";
  const resourceGroupName = "cs-tester";
  const cloudServiceName = "cs1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.publicIPAddresses.listCloudServicePublicIPAddresses(
    resourceGroupName,
    cloudServiceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listVmssPublicIP().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.
