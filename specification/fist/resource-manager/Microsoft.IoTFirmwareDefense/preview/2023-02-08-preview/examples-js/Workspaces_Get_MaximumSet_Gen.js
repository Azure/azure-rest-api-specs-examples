const { IoTFirmwareDefenseClient } = require("@azure/arm-iotfirmwaredefense");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get firmware analysis workspace.
 *
 * @summary Get firmware analysis workspace.
 * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Workspaces_Get_MaximumSet_Gen.json
 */
async function workspacesGetMaximumSetGen() {
  const subscriptionId =
    process.env["IOTFIRMWAREDEFENSE_SUBSCRIPTION_ID"] || "9781B4B5-0922-472A-80F0-B743D0596694";
  const resourceGroupName = process.env["IOTFIRMWAREDEFENSE_RESOURCE_GROUP"] || "rgworkspaces";
  const workspaceName = "E_US";
  const credential = new DefaultAzureCredential();
  const client = new IoTFirmwareDefenseClient(credential, subscriptionId);
  const result = await client.workspaces.get(resourceGroupName, workspaceName);
  console.log(result);
}
