const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of Network Experiment Profiles within a resource group under a subscription
 *
 * @summary Gets a list of Network Experiment Profiles within a resource group under a subscription
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentListProfiles.json
 */
async function listNetworkExperimentProfilesInAResourceGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "MyResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.networkExperimentProfiles.listByResourceGroup(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listNetworkExperimentProfilesInAResourceGroup().catch(console.error);
