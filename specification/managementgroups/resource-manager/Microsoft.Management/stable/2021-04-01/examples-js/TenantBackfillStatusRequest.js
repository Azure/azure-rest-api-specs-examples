const { ManagementGroupsAPI } = require("@azure/arm-managementgroups");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets tenant backfill status
 *
 * @summary Gets tenant backfill status
 * x-ms-original-file: specification/managementgroups/resource-manager/Microsoft.Management/stable/2021-04-01/examples/TenantBackfillStatusRequest.json
 */
async function tenantBackfillStatus() {
  const credential = new DefaultAzureCredential();
  const client = new ManagementGroupsAPI(credential);
  const result = await client.tenantBackfillStatus();
  console.log(result);
}

tenantBackfillStatus().catch(console.error);
