const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the details of a specified job on a Data Box Edge/Data Box Gateway device.
 *
 * @summary Gets the details of a specified job on a Data Box Edge/Data Box Gateway device.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/OperationsStatusGet.json
 */
async function operationsStatusGet() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const name = "159a00c7-8543-4343-9435-263ac87df3bb";
  const resourceGroupName = "GroupForEdgeAutomation";
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const result = await client.operationsStatus.get(deviceName, name, resourceGroupName);
  console.log(result);
}

operationsStatusGet().catch(console.error);
