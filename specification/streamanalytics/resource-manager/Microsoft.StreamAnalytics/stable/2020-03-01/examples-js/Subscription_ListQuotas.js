const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieves the subscription's current quota information in a particular region.
 *
 * @summary Retrieves the subscription's current quota information in a particular region.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Subscription_ListQuotas.json
 */
async function listSubscriptionQuotaInformationInWestUs() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const location = "West US";
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.subscriptions.listQuotas(location);
  console.log(result);
}

listSubscriptionQuotaInformationInWestUs().catch(console.error);
