const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the data connector.
 *
 * @summary creates or updates the data connector.
 * x-ms-original-file: 2025-07-01-preview/dataConnectors/CreateDynamics365DataConnetor.json
 */
async function createsOrUpdatesADynamics365DataConnector() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.dataConnectors.createOrUpdate(
    "myRg",
    "myWorkspace",
    "c2541efb-c9a6-47fe-9501-87d1017d1512",
    {
      etag: '"0300bf09-0000-0000-0000-5c37296e0000"',
      kind: "Dynamics365",
      dataTypes: { dynamics365CdsActivities: { state: "Enabled" } },
      tenantId: "2070ecc9-b4d5-4ae4-adaa-936fa1954fa8",
    },
  );
  console.log(result);
}
