const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get a list of all tenant action groups in a management group.
 *
 * @summary Get a list of all tenant action groups in a management group.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/preview/2023-05-01-preview/examples/listTenantActionGroups.json
 */
async function listTenantActionGroupsAtManagementGroupLevel() {
  const managementGroupId = "72f988bf-86f1-41af-91ab-2d7cd011db47";
  const xMsClientTenantId = "72f988bf-86f1-41af-91ab-2d7cd011db47";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential);
  const resArray = new Array();
  for await (let item of client.tenantActionGroups.listByManagementGroupId(
    managementGroupId,
    xMsClientTenantId,
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
