const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all application security groups in a subscription.
 *
 * @summary Gets all application security groups in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/ApplicationSecurityGroupListAll.json
 */
async function listAllApplicationSecurityGroups() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.applicationSecurityGroups.listAll()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllApplicationSecurityGroups().catch(console.error);
