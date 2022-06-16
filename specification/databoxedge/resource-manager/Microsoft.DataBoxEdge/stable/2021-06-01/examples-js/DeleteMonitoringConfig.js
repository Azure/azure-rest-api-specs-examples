const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to deletes a new metric configuration for a role.
 *
 * @summary deletes a new metric configuration for a role.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/DeleteMonitoringConfig.json
 */
async function deleteMonitoringConfig() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const roleName = "testrole";
  const resourceGroupName = "GroupForEdgeAutomation";
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const result = await client.monitoringConfig.beginDeleteAndWait(
    deviceName,
    roleName,
    resourceGroupName
  );
  console.log(result);
}

deleteMonitoringConfig().catch(console.error);
