const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete a tenant action group.
 *
 * @summary Delete a tenant action group.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2023-05-01-preview/examples/deleteTenantActionGroup.json
 */
async function deleteATenantActionGroup() {
  const managementGroupId = "72f988bf-86f1-41af-91ab-2d7cd011db47";
  const tenantActionGroupName = "testTenantActionGroup";
  const xMsClientTenantId = "72f988bf-86f1-41af-91ab-2d7cd011db47";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const result = await client.tenantActionGroups.delete(
    managementGroupId,
    tenantActionGroupName,
    xMsClientTenantId,
  );
  console.log(result);
}
