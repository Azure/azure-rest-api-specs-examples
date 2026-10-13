const { ContentStoreClient } = require("@azure/arm-commvaultcontentstore");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to resolve the token to get the SaaS resource ID and activate the SaaS resource
 *
 * @summary resolve the token to get the SaaS resource ID and activate the SaaS resource
 * x-ms-original-file: 2026-09-30/SaaSOperationGroup_ActivateResource_MinimumSet_Gen.json
 */
async function saaSOperationGroupActivateResourceMaximumSetGeneratedByMaximumSetRuleGeneratedByMinimumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ContentStoreClient(credential, subscriptionId);
  const result = await client.saaSOperationGroup.activateResource({
    saasGuid: "55555555-6666-7777-8888-999999999999",
  });
  console.log(result);
}
