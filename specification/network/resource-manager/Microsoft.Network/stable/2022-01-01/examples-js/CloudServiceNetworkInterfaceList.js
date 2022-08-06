const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all network interfaces in a cloud service.
 *
 * @summary Gets all network interfaces in a cloud service.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/CloudServiceNetworkInterfaceList.json
 */
async function listCloudServiceNetworkInterfaces() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const cloudServiceName = "cs1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.networkInterfaces.listCloudServiceNetworkInterfaces(
    resourceGroupName,
    cloudServiceName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listCloudServiceNetworkInterfaces().catch(console.error);
