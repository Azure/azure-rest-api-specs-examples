const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Update linked storage accounts for an Application Insights component.
 *
 * @summary Update linked storage accounts for an Application Insights component.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2020-03-01-preview/examples/ComponentLinkedStorageAccountsUpdate.json
 */
async function componentLinkedStorageAccountsUpdate() {
  const subscriptionId =
    process.env["APPLICATIONINSIGHTS_SUBSCRIPTION_ID"] || "86dc51d3-92ed-4d7e-947a-775ea79b4918";
  const resourceGroupName =
    process.env["APPLICATIONINSIGHTS_RESOURCE_GROUP"] || "someResourceGroupName";
  const resourceName = "myComponent";
  const storageType = "ServiceProfiler";
  const linkedStorageAccountsProperties = {
    linkedStorageAccount:
      "/subscriptions/86dc51d3-92ed-4d7e-947a-775ea79b4918/resourceGroups/someResourceGroupName/providers/Microsoft.Storage/storageAccounts/storageaccountname",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.componentLinkedStorageAccountsOperations.update(
    resourceGroupName,
    resourceName,
    storageType,
    linkedStorageAccountsProperties,
  );
  console.log(result);
}
