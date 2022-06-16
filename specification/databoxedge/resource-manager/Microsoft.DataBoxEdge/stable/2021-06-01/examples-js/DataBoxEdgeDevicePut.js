const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Data Box Edge/Data Box Gateway resource.
 *
 * @summary Creates or updates a Data Box Edge/Data Box Gateway resource.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/DataBoxEdgeDevicePut.json
 */
async function dataBoxEdgeDevicePut() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const resourceGroupName = "GroupForEdgeAutomation";
  const dataBoxEdgeDevice = {
    location: "WUS",
    sku: { name: "Edge", tier: "Standard" },
    tags: {},
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const result = await client.devices.createOrUpdate(
    deviceName,
    resourceGroupName,
    dataBoxEdgeDevice
  );
  console.log(result);
}

dataBoxEdgeDevicePut().catch(console.error);
