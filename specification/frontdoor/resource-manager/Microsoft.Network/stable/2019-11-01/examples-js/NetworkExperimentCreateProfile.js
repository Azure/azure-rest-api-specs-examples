const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an NetworkExperiment Profile
 *
 * @summary Creates an NetworkExperiment Profile
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentCreateProfile.json
 */
async function createsAnNetworkExperimentProfileInAResourceGroup() {
  const subscriptionId = "subid";
  const profileName = "MyProfile";
  const resourceGroupName = "MyResourceGroup";
  const parameters = { enabledState: "Enabled", location: "WestUs" };
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const result = await client.networkExperimentProfiles.beginCreateOrUpdateAndWait(
    profileName,
    resourceGroupName,
    parameters
  );
  console.log(result);
}

createsAnNetworkExperimentProfileInAResourceGroup().catch(console.error);
