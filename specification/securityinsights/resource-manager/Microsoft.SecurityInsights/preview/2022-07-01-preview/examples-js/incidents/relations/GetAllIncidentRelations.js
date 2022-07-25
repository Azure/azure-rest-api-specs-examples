const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets all incident relations.
 *
 * @summary Gets all incident relations.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/incidents/relations/GetAllIncidentRelations.json
 */
async function getAllIncidentRelations() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const incidentId = "afbd324f-6c48-459c-8710-8d1e1cd03812";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.incidentRelations.list(
    resourceGroupName,
    workspaceName,
    incidentId
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

getAllIncidentRelations().catch(console.error);
