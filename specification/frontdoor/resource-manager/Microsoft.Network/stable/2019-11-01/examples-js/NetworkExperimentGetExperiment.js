const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an Experiment by ExperimentName
 *
 * @summary Gets an Experiment by ExperimentName
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentGetExperiment.json
 */
async function getsAnExperimentByExperimentName() {
  const subscriptionId = "subid";
  const resourceGroupName = "MyResourceGroup";
  const profileName = "MyProfile";
  const experimentName = "MyExperiment";
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const result = await client.experiments.get(resourceGroupName, profileName, experimentName);
  console.log(result);
}

getsAnExperimentByExperimentName().catch(console.error);
