const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create a Metadata.
 *
 * @summary create a Metadata.
 * x-ms-original-file: 2025-07-01-preview/metadata/PutMetadataMinimal.json
 */
async function createOrUpdateMinimalMetadata() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.metadata.create("myRg", "myWorkspace", "metadataName", {
    contentId: "c00ee137-7475-47c8-9cce-ec6f0f1bedd0",
    kind: "AnalyticsRule",
    parentId:
      "/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName",
  });
  console.log(result);
}
