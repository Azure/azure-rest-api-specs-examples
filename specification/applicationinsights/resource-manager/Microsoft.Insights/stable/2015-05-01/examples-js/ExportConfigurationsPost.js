const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Create a Continuous Export configuration of an Application Insights component.
 *
 * @summary Create a Continuous Export configuration of an Application Insights component.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/ExportConfigurationsPost.json
 */
async function exportConfigurationPost() {
  const subscriptionId = process.env["APPLICATIONINSIGHTS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName =
    process.env["APPLICATIONINSIGHTS_RESOURCE_GROUP"] || "my-resource-group";
  const resourceName = "my-component";
  const exportProperties = {
    destinationAccountId:
      "/subscriptions/subid/resourceGroups/my-resource-group/providers/Microsoft.ClassicStorage/storageAccounts/mystorageblob",
    destinationAddress:
      "https://mystorageblob.blob.core.windows.net/testexport?sv=2015-04-05&sr=c&sig=token",
    destinationStorageLocationId: "eastus",
    destinationStorageSubscriptionId: "subid",
    destinationType: "Blob",
    isEnabled: "true",
    notificationQueueEnabled: "false",
    notificationQueueUri: "",
    recordTypes:
      "Requests, Event, Exceptions, Metrics, PageViews, PageViewPerformance, Rdd, PerformanceCounters, Availability",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.exportConfigurations.create(
    resourceGroupName,
    resourceName,
    exportProperties,
  );
  console.log(result);
}
