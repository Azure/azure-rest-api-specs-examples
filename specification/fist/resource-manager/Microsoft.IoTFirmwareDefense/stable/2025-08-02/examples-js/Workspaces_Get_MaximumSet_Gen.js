const { IoTFirmwareDefenseClient } = require("@azure/arm-iotfirmwaredefense");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to get firmware analysis workspace.
 *
 * @summary get firmware analysis workspace.
 * x-ms-original-file: 2025-08-02/Workspaces_Get_MaximumSet_Gen.json
 */
async function workspacesGetMaximumSetGenGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "5C707B5F-6130-4F71-819E-953A28942E88";
  const client = new IoTFirmwareDefenseClient(credential, subscriptionId);
  const result = await client.workspaces.get("rgiotfirmwaredefense", "exampleWorkspaceName");
  console.log(result);
}
