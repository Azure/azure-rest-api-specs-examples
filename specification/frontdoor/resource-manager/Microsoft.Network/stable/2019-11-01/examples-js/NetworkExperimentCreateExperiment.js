const { FrontDoorManagementClient } = require("@azure/arm-frontdoor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an Experiment
 *
 * @summary Creates or updates an Experiment
 * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2019-11-01/examples/NetworkExperimentCreateExperiment.json
 */
async function createsAnExperiment() {
  const subscriptionId = process.env["FRONTDOOR_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["FRONTDOOR_RESOURCE_GROUP"] || "MyResourceGroup";
  const profileName = "MyProfile";
  const experimentName = "MyExperiment";
  const parameters = {
    description: "this is my first experiment!",
    enabledState: "Enabled",
    endpointA: { name: "endpoint A", endpoint: "endpointA.net" },
    endpointB: { name: "endpoint B", endpoint: "endpointB.net" },
  };
  const credential = new DefaultAzureCredential();
  const client = new FrontDoorManagementClient(credential, subscriptionId);
  const result = await client.experiments.beginCreateOrUpdateAndWait(
    resourceGroupName,
    profileName,
    experimentName,
    parameters,
  );
  console.log(result);
}
