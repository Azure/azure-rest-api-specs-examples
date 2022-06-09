```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about all network interfaces in a role instance in a cloud service.
 *
 * @summary Gets information about all network interfaces in a role instance in a cloud service.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/CloudServiceRoleInstanceNetworkInterfaceList.json
 */
async function listCloudServiceRoleInstanceNetworkInterfaces() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const cloudServiceName = "cs1";
  const roleInstanceName = "TestVMRole_IN_0";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.networkInterfaces.listCloudServiceRoleInstanceNetworkInterfaces(
    resourceGroupName,
    cloudServiceName,
    roleInstanceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listCloudServiceRoleInstanceNetworkInterfaces().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_28.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.
