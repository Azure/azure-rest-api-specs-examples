const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Return trial status for subscription by region
 *
 * @summary Return trial status for subscription by region
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/Locations_CheckTrialAvailabilityWithSku.json
 */
async function locationsCheckTrialAvailabilityWithSku() {
  const subscriptionId =
    process.env["AVS_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const location = "eastus";
  const sku = { name: "avs52t" };
  const options = { sku };
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const result = await client.locations.checkTrialAvailability(location, options);
  console.log(result);
}
