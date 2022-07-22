const { BatchManagementClient } = require("@azure/arm-batch");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks whether the Batch account name is available in the specified region.
 *
 * @summary Checks whether the Batch account name is available in the specified region.
 * x-ms-original-file: specification/batch/resource-manager/Microsoft.Batch/stable/2022-06-01/examples/LocationCheckNameAvailability_AlreadyExists.json
 */
async function locationCheckNameAvailabilityAlreadyExists() {
  const subscriptionId = "subid";
  const locationName = "japaneast";
  const parameters = {
    name: "existingaccountname",
    type: "Microsoft.Batch/batchAccounts",
  };
  const credential = new DefaultAzureCredential();
  const client = new BatchManagementClient(credential, subscriptionId);
  const result = await client.location.checkNameAvailability(locationName, parameters);
  console.log(result);
}

locationCheckNameAvailabilityAlreadyExists().catch(console.error);
