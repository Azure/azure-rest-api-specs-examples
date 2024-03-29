const { BillingManagementClient } = require("@azure/arm-billing");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a named billing period.  This is only supported for Azure Web-Direct subscriptions. Other subscription types which were not purchased directly through the Azure web portal are not supported through this preview API.
 *
 * @summary Gets a named billing period.  This is only supported for Azure Web-Direct subscriptions. Other subscription types which were not purchased directly through the Azure web portal are not supported through this preview API.
 * x-ms-original-file: specification/billing/resource-manager/Microsoft.Billing/preview/2018-03-01-preview/examples/BillingPeriodsGet.json
 */
async function billingPeriodsGet() {
  const subscriptionId = "subid";
  const billingPeriodName = "201702-1";
  const credential = new DefaultAzureCredential();
  const client = new BillingManagementClient(credential, subscriptionId);
  const result = await client.billingPeriods.get(billingPeriodName);
  console.log(result);
}

billingPeriodsGet().catch(console.error);
