const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List allowed upgrade plans for application.
 *
 * @summary List allowed upgrade plans for application.
 * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/listAllowedUpgradePlans.json
 */
async function listAllowedUpgradePlansForApplication() {
  const subscriptionId = process.env["MANAGEDAPPLICATIONS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["MANAGEDAPPLICATIONS_RESOURCE_GROUP"] || "rg";
  const applicationName = "myManagedApplication";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential, subscriptionId);
  const result = await client.applications.listAllowedUpgradePlans(
    resourceGroupName,
    applicationName
  );
  console.log(result);
}
