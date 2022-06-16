const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of application control machine groups for the subscription.
 *
 * @summary Gets a list of application control machine groups for the subscription.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2020-01-01/examples/ApplicationWhitelistings/GetAdaptiveApplicationControlsSubscription_example.json
 */
async function getsAListOfApplicationControlGroupsOfMachinesForTheSubscription() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const includePathRecommendations = true;
  const summary = false;
  const options = {
    includePathRecommendations,
    summary,
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.adaptiveApplicationControls.list(options);
  console.log(result);
}

getsAListOfApplicationControlGroupsOfMachinesForTheSubscription().catch(console.error);
