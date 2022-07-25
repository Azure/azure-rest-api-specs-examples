const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a Metadata.
 *
 * @summary Create a Metadata.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/metadata/PutMetadataMinimal.json
 */
async function createOrUpdateMinimalMetadata() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const metadataName = "metadataName";
  const metadata = {
    contentId: "c00ee137-7475-47c8-9cce-ec6f0f1bedd0",
    kind: "AnalyticsRule",
    parentId:
      "/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName",
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.metadata.create(
    resourceGroupName,
    workspaceName,
    metadataName,
    metadata
  );
  console.log(result);
}

createOrUpdateMinimalMetadata().catch(console.error);
