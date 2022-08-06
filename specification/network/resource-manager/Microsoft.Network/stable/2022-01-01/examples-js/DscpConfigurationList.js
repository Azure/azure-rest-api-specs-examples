const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a DSCP Configuration.
 *
 * @summary Gets a DSCP Configuration.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/DscpConfigurationList.json
 */
async function getDscpConfiguration() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.dscpConfigurationOperations.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

getDscpConfiguration().catch(console.error);
