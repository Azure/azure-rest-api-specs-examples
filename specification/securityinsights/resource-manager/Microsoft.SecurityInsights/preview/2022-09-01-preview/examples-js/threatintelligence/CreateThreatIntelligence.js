const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create a new threat intelligence indicator.
 *
 * @summary Create a new threat intelligence indicator.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/threatintelligence/CreateThreatIntelligence.json
 */
async function createANewThreatIntelligence() {
  const subscriptionId =
    process.env["SECURITYINSIGHT_SUBSCRIPTION_ID"] || "bd794837-4d29-4647-9105-6339bfdb4e6a";
  const resourceGroupName = process.env["SECURITYINSIGHT_RESOURCE_GROUP"] || "myRg";
  const workspaceName = "myWorkspace";
  const threatIntelligenceProperties = {
    description: "debugging indicators",
    confidence: 78,
    createdByRef: "contoso@contoso.com",
    displayName: "new schema",
    externalReferences: [],
    granularMarkings: [],
    killChainPhases: [],
    kind: "indicator",
    labels: [],
    modified: "",
    pattern: "[url:value = 'https://www.contoso.com']",
    patternType: "url",
    revoked: false,
    source: "Azure Sentinel",
    threatIntelligenceTags: ["new schema"],
    threatTypes: ["compromised"],
    validFrom: "2021-09-15T17:44:00.114052Z",
    validUntil: "",
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.threatIntelligenceIndicator.createIndicator(
    resourceGroupName,
    workspaceName,
    threatIntelligenceProperties
  );
  console.log(result);
}
