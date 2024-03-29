const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the available billing periods for a subscription in reverse chronological order. This is only supported for Azure Web-Direct subscriptions. Other subscription types which were not purchased directly through the Azure web portal are not supported through this preview API.
 *
 * @summary Lists the available billing periods for a subscription in reverse chronological order. This is only supported for Azure Web-Direct subscriptions. Other subscription types which were not purchased directly through the Azure web portal are not supported through this preview API.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/preview/2018-03-01-preview/examples/BillingPeriodsList.json
 */
async function billingPeriodsList() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.billingPeriods.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

billingPeriodsList().catch(console.error);
