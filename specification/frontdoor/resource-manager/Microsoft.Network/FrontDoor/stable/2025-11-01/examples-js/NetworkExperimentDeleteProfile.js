const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to deletes an NetworkExperiment Profile by ProfileName
 *
 * @summary deletes an NetworkExperiment Profile by ProfileName
 * x-ms-original-file: 2025-11-01/NetworkExperimentDeleteProfile.json
 */
async function deletesAnNetworkExperimentProfileByProfileName() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subid";
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  await client.networkExperimentProfiles.delete("MyResourceGroup", "MyProfile");
}
