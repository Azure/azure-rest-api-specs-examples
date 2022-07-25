const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update a threat Intelligence indicator.
 *
 * @summary Update a threat Intelligence indicator.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/threatintelligence/UpdateThreatIntelligence.json
 */
async function updateAThreatIntelligenceIndicator() {
  const subscriptionId = "bd794837-4d29-4647-9105-6339bfdb4e6a";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const name = "d9cd6f0b-96b9-3984-17cd-a779d1e15a93";
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
    validFrom: "2020-04-15T17:44:00.114052Z",
    validUntil: "",
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.threatIntelligenceIndicator.create(
    resourceGroupName,
    workspaceName,
    name,
    threatIntelligenceProperties
  );
  console.log(result);
}

updateAThreatIntelligenceIndicator().catch(console.error);
