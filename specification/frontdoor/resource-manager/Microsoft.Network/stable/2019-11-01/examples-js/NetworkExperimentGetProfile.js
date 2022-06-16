const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an NetworkExperiment Profile by ProfileName
 *
 * @summary Gets an NetworkExperiment Profile by ProfileName
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentGetProfile.json
 */
async function getsAnNetworkExperimentProfileByProfileId() {
  const subscriptionId = "subid";
  const resourceGroupName = "MyResourceGroup";
  const profileName = "MyProfile";
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const result = await client.networkExperimentProfiles.get(resourceGroupName, profileName);
  console.log(result);
}

getsAnNetworkExperimentProfileByProfileId().catch(console.error);
