const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all the SIM policies in a mobile network.
 *
 * @summary Gets all the SIM policies in a mobile network.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SimPolicyListByMobileNetwork.json
 */
async function listSimPoliciesInAMobileNetwork() {
  const subscriptionId = "subid";
  const resourceGroupName = "testResourceGroupName";
  const mobileNetworkName = "testMobileNetwork";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.simPolicies.listByMobileNetwork(
    resourceGroupName,
    mobileNetworkName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listSimPoliciesInAMobileNetwork().catch(console.error);
