const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about all public IP addresses in a role instance IP configuration in a cloud service.
 *
 * @summary Gets information about all public IP addresses in a role instance IP configuration in a cloud service.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/CloudServiceRoleInstancePublicIpList.json
 */
async function listVmssvmPublicIP() {
  const subscriptionId = "subid";
  const resourceGroupName = "cs-tester";
  const cloudServiceName = "cs1";
  const roleInstanceName = "Test_VM_0";
  const networkInterfaceName = "nic1";
  const ipConfigurationName = "ip1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.publicIPAddresses.listCloudServiceRoleInstancePublicIPAddresses(
    resourceGroupName,
    cloudServiceName,
    roleInstanceName,
    networkInterfaceName,
    ipConfigurationName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listVmssvmPublicIP().catch(console.error);
