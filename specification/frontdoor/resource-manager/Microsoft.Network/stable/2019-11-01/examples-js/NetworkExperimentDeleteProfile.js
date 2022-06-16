const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an NetworkExperiment Profile by ProfileName
 *
 * @summary Deletes an NetworkExperiment Profile by ProfileName
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentDeleteProfile.json
 */
async function deletesAnNetworkExperimentProfileByProfileName() {
  const subscriptionId = "subid";
  const resourceGroupName = "MyResourceGroup";
  const profileName = "MyProfile";
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const result = await client.networkExperimentProfiles.beginDeleteAndWait(
    resourceGroupName,
    profileName
  );
  console.log(result);
}

deletesAnNetworkExperimentProfileByProfileName().catch(console.error);
