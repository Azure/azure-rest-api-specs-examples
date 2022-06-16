const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the managed application.
 *
 * @summary Deletes the managed application.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/deleteApplication.json
 */
async function deletesAManagedApplication() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg";
  const applicationName = "myManagedApplication";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential, subscriptionId);
  const result = await client.applications.beginDeleteAndWait(resourceGroupName, applicationName);
  console.log(result);
}

deletesAManagedApplication().catch(console.error);
