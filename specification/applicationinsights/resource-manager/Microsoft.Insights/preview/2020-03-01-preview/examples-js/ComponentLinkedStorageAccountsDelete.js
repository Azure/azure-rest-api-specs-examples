const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete linked storage accounts for an Application Insights component.
 *
 * @summary Delete linked storage accounts for an Application Insights component.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/preview/2020-03-01-preview/examples/ComponentLinkedStorageAccountsDelete.json
 */
async function componentLinkedStorageAccountsDelete() {
  const subscriptionId = "86dc51d3-92ed-4d7e-947a-775ea79b4918";
  const resourceGroupName = "someResourceGroupName";
  const resourceName = "myComponent";
  const storageType = "ServiceProfiler";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.componentLinkedStorageAccountsOperations.delete(
    resourceGroupName,
    resourceName,
    storageType
  );
  console.log(result);
}

componentLinkedStorageAccountsDelete().catch(console.error);
