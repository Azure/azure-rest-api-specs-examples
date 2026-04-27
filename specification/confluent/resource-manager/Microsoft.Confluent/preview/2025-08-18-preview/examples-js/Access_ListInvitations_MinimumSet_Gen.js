const { ConfluentManagementClient } = require("@azure/arm-confluent");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to organization accounts invitation details
 *
 * @summary organization accounts invitation details
 * x-ms-original-file: 2025-08-18-preview/Access_ListInvitations_MinimumSet_Gen.json
 */
async function accessListInvitationsMinimumSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "DC34558A-05D3-4370-AED8-75E60B381F94";
  const client = new ConfluentManagementClient(credential, subscriptionId);
  const result = await client.access.listInvitations(
    "rgconfluent",
    "edpxevovxieanzlscvflmmcuoracwh",
    {},
  );
  console.log(result);
}
