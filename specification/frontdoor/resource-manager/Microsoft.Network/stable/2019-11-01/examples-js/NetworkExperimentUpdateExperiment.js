const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an Experiment
 *
 * @summary Updates an Experiment
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentUpdateExperiment.json
 */
async function updatesAnExperiment() {
  const subscriptionId = "subid";
  const resourceGroupName = "MyResourceGroup";
  const profileName = "MyProfile";
  const experimentName = "MyExperiment";
  const parameters = {
    description: "string",
    enabledState: "Enabled",
  };
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const result = await client.experiments.beginUpdateAndWait(
    resourceGroupName,
    profileName,
    experimentName,
    parameters
  );
  console.log(result);
}

updatesAnExperiment().catch(console.error);
