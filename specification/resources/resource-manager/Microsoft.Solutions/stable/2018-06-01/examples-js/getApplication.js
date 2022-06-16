const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the managed application.
 *
 * @summary Gets the managed application.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/getApplication.json
 */
async function getAManagedApplication() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg";
  const applicationName = "myManagedApplication";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential, subscriptionId);
  const result = await client.applications.get(resourceGroupName, applicationName);
  console.log(result);
}

getAManagedApplication().catch(console.error);
