const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks whether the Batch account name is available in the specified region.
 *
 * @summary Checks whether the Batch account name is available in the specified region.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-01-01/examples/LocationCheckNameAvailability_Available.json
 */
async function locationCheckNameAvailabilityAvailable() {
  const subscriptionId = "subid";
  const locationName = "japaneast";
  const parameters = {
    name: "newaccountname",
    type: "Microsoft.Batch/batchAccounts",
  };
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.location.checkNameAvailability(locationName, parameters);
  console.log(result);
}

locationCheckNameAvailabilityAvailable().catch(console.error);
