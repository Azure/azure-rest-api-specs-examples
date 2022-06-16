const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the managed application.
 *
 * @summary Gets the managed application.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/getApplicationById.json
 */
async function getApplicationById() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const applicationId = "myApplicationId";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential, subscriptionId);
  const result = await client.applications.getById(applicationId);
  console.log(result);
}

getApplicationById().catch(console.error);
