const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the data connector.
 *
 * @summary Creates or updates the data connector.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/dataConnectors/CreateDynamics365DataConnetor.json
 */
async function createsOrUpdatesADynamics365DataConnector() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const dataConnectorId = "c2541efb-c9a6-47fe-9501-87d1017d1512";
  const dataConnector = {
    dataTypes: { dynamics365CdsActivities: { state: "Enabled" } },
    etag: '"0300bf09-0000-0000-0000-5c37296e0000"',
    kind: "Dynamics365",
    tenantId: "2070ecc9-b4d5-4ae4-adaa-936fa1954fa8",
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.dataConnectors.createOrUpdate(
    resourceGroupName,
    workspaceName,
    dataConnectorId,
    dataConnector
  );
  console.log(result);
}

createsOrUpdatesADynamics365DataConnector().catch(console.error);
