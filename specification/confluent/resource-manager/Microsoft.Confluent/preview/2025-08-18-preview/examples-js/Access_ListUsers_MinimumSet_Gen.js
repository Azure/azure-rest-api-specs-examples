const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to organization users details
 *
 * @summary organization users details
 * x-ms-original-file: 2025-08-18-preview/Access_ListUsers_MinimumSet_Gen.json
 */
async function accessListUsersMinimumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "DC34558A-05D3-4370-AED8-75E60B381F94";
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const result = await client.access.listUsers("rgconfluent", "elqetgujssclojggilbgl", {});
  console.log(result);
}
