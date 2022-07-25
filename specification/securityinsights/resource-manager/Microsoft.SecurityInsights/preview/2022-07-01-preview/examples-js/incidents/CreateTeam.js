const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a Microsoft team to investigate the incident by sharing information and insights between participants.
 *
 * @summary Creates a Microsoft team to investigate the incident by sharing information and insights between participants.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/incidents/CreateTeam.json
 */
async function createsIncidentTeamsGroup() {
  const subscriptionId = "9023f5b5-df22-4313-8fbf-b4b75af8a6d9";
  const resourceGroupName = "ambawolvese5resourcegroup";
  const workspaceName = "AmbaE5WestCentralUS";
  const incidentId = "69a30280-6a4c-4aa7-9af0-5d63f335d600";
  const teamProperties = {
    teamDescription: "Team description",
    teamName: "Team name",
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.incidents.createTeam(
    resourceGroupName,
    workspaceName,
    incidentId,
    teamProperties
  );
  console.log(result);
}

createsIncidentTeamsGroup().catch(console.error);
