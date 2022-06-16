const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Modifies a Data Box Edge/Data Box Gateway resource.
 *
 * @summary Modifies a Data Box Edge/Data Box Gateway resource.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/DataBoxEdgeDevicePatch.json
 */
async function dataBoxEdgeDevicePatch() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const resourceGroupName = "GroupForEdgeAutomation";
  const parameters = {
    edgeProfile: {
      subscription: {
        id: "/subscriptions/0d44739e-0563-474f-97e7-24a0cdb23b29/resourceGroups/rapvs-rg/providers/Microsoft.AzureStack/linkedSubscriptions/ca014ddc-5cf2-45f8-b390-e901e4a0ae87",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const result = await client.devices.update(deviceName, resourceGroupName, parameters);
  console.log(result);
}

dataBoxEdgeDevicePatch().catch(console.error);
