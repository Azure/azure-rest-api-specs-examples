const { PlatformValidationClient } = require("@azure/arm-platformvalidation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list validation test category catalog entries for a subscription
 *
 * @summary list validation test category catalog entries for a subscription
 * x-ms-original-file: 2026-08-01-preview/ValidationTestCategories_ListBySubscription_MaximumSet_Gen.json
 */
async function validationTestCategoriesListBySubscriptionMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new PlatformValidationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.validationTestCategories.listBySubscription({
    filter: "audience eq 'Public'",
  })) {
    resArray.push(item);
  }

  console.log(resArray);
}
