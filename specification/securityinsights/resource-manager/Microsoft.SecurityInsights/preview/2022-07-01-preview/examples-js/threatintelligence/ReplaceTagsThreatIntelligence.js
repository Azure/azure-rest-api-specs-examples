const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Replace tags added to a threat intelligence indicator.
 *
 * @summary Replace tags added to a threat intelligence indicator.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/threatintelligence/ReplaceTagsThreatIntelligence.json
 */
async function replaceTagsToAThreatIntelligence() {
  const subscriptionId = "bd794837-4d29-4647-9105-6339bfdb4e6a";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const name = "d9cd6f0b-96b9-3984-17cd-a779d1e15a93";
  const threatIntelligenceReplaceTags = {
    etag: '"0000262c-0000-0800-0000-5e9767060000"',
    kind: "indicator",
    threatIntelligenceTags: ["patching tags"],
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.threatIntelligenceIndicator.replaceTags(
    resourceGroupName,
    workspaceName,
    name,
    threatIntelligenceReplaceTags
  );
  console.log(result);
}

replaceTagsToAThreatIntelligence().catch(console.error);
