const { PlatformValidationClient } = require("@azure/arm-platformvalidation");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list validation test catalog entries for a subscription
 *
 * @summary list validation test catalog entries for a subscription
 * x-ms-original-file: 2026-08-01-preview/ValidationTests_ListBySubscription_MaximumSet_Gen.json
 */
async function validationTestsListBySubscriptionMaximumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new PlatformValidationClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.validationTests.listBySubscription({
    filter: "audience eq 'Public'",
  })) {
    resArray.push(item);
  }

  console.log(resArray);
}
