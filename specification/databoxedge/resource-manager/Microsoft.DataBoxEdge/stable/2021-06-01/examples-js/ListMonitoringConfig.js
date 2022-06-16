const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists metric configurations in a role.
 *
 * @summary Lists metric configurations in a role.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/ListMonitoringConfig.json
 */
async function listMonitoringConfig() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const roleName = "testrole";
  const resourceGroupName = "GroupForEdgeAutomation";
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.monitoringConfig.list(deviceName, roleName, resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listMonitoringConfig().catch(console.error);
