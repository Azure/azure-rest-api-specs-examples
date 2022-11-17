const { ApplicationInsightsManagementClient } = require("@azure/arm-appinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update current billing features for an Application Insights component.
 *
 * @summary Update current billing features for an Application Insights component.
 * x-ms-original-file: specification/applicationinsights/resource-manager/Microsoft.Insights/stable/2015-05-01/examples/CurrentBillingFeaturesUpdate.json
 */
async function componentCurrentBillingFeaturesUpdate() {
  const subscriptionId = "subid";
  const resourceGroupName = "my-resource-group";
  const resourceName = "my-component";
  const billingFeaturesProperties = {
    currentBillingFeatures: ["Basic", "Application Insights Enterprise"],
    dataVolumeCap: { cap: 100, stopSendNotificationWhenHitCap: true },
  };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationInsightsManagementClient(credential, subscriptionId);
  const result = await client.componentCurrentBillingFeatures.update(
    resourceGroupName,
    resourceName,
    billingFeaturesProperties
  );
  console.log(result);
}

componentCurrentBillingFeaturesUpdate().catch(console.error);
