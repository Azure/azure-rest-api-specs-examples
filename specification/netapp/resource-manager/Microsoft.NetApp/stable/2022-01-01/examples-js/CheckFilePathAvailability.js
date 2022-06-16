const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Check if a file path is available.
 *
 * @summary Check if a file path is available.
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-01-01/examples/CheckFilePathAvailability.json
 */
async function checkFilePathAvailability() {
  const subscriptionId = "D633CC2E-722B-4AE1-B636-BBD9E4C60ED9";
  const location = "eastus";
  const name = "my-exact-filepth";
  const subnetId =
    "/subscriptions/9760acf5-4638-11e7-9bdb-020073ca7778/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.netAppResource.checkFilePathAvailability(location, name, subnetId);
  console.log(result);
}

checkFilePathAvailability().catch(console.error);
