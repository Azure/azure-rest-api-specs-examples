const { IoTFirmwareDefenseClient } = require("@azure/arm-iotfirmwaredefense");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get firmware analysis workspace.
 *
 * @summary get firmware analysis workspace.
 * x-ms-original-file: 2025-08-02/Workspaces_Get_MinimumSet_Gen.json
 */
async function workspacesGetMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new IoTFirmwareDefenseClient(credential, subscriptionId);
  const result = await client.workspaces.get("rgworkspaces", "E_US");
  console.log(result);
}
