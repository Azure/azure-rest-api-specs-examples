const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the Batch service quotas for the specified subscription at the given location.
 *
 * @summary Gets the Batch service quotas for the specified subscription at the given location.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-06-01/examples/LocationGetQuotas.json
 */
async function locationGetQuotas() {
  const subscriptionId = "subid";
  const locationName = "japaneast";
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.location.getQuotas(locationName);
  console.log(result);
}

locationGetQuotas().catch(console.error);
