const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a addon.
 *
 * @summary Create or update a addon.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/PutAddons.json
 */
async function putAddOns() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const roleName = "KubernetesRole";
  const addonName = "arcName";
  const resourceGroupName = "GroupForEdgeAutomation";
  const addon = {
    kind: "ArcForKubernetes",
    resourceGroupName: "GroupForEdgeAutomation",
    resourceLocation: "EastUS",
    resourceName: "testedgedevice",
    subscriptionId: "4385cf00-2d3a-425a-832f-f4285b1c9dce",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const result = await client.addons.beginCreateOrUpdateAndWait(
    deviceName,
    roleName,
    addonName,
    resourceGroupName,
    addon
  );
  console.log(result);
}

putAddOns().catch(console.error);
