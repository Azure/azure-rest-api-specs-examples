const { AtlasClient } = require("@azure/arm-mongodbatlas");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to list OrganizationResource resources by subscription ID
 *
 * @summary list OrganizationResource resources by subscription ID
 * x-ms-original-file: 2024-11-18-preview/Organizations_ListBySubscription_MinimumSet_Gen.json
 */
async function organizationsListBySubscriptionMaximumSetGeneratedByMinimumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "422A4D59-A5BC-4DBB-8831-EC666633F64F";
  const client = new AtlasClient(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.organizations.listBySubscription()) {
    resArray.push(item);
  }

  console.log(resArray);
}
