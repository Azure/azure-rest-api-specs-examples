const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Query threat intelligence indicators as per filtering criteria.
 *
 * @summary Query threat intelligence indicators as per filtering criteria.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/threatintelligence/QueryThreatIntelligence.json
 */
async function queryThreatIntelligenceIndicatorsAsPerFilteringCriteria() {
  const subscriptionId =
    process.env["SECURITYINSIGHT_SUBSCRIPTION_ID"] || "bd794837-4d29-4647-9105-6339bfdb4e6a";
  const resourceGroupName = process.env["SECURITYINSIGHT_RESOURCE_GROUP"] || "myRg";
  const workspaceName = "myWorkspace";
  const threatIntelligenceFilteringCriteria = {
    maxConfidence: 80,
    maxValidUntil: "2021-04-25T17:44:00.114052Z",
    minConfidence: 25,
    minValidUntil: "2021-04-05T17:44:00.114052Z",
    pageSize: 100,
    sortBy: [{ itemKey: "lastUpdatedTimeUtc", sortOrder: "descending" }],
    sources: ["Azure Sentinel"],
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.threatIntelligenceIndicator.listQueryIndicators(
    resourceGroupName,
    workspaceName,
    threatIntelligenceFilteringCriteria
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
