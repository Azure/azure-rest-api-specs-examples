const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to return trial status for subscription by region
 *
 * @summary return trial status for subscription by region
 * x-ms-original-file: 2025-09-01/Locations_CheckTrialAvailability.json
 */
async function locationsCheckTrialAvailability() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.locations.checkTrialAvailability("eastus");
  console.log(result);
}
