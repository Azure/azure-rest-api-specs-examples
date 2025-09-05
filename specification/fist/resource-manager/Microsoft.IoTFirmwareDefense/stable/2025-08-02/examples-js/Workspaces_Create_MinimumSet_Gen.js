const { IoTFirmwareDefenseClient } = require("@azure/arm-iotfirmwaredefense");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a firmware analysis workspace.
 *
 * @summary the operation to create or update a firmware analysis workspace.
 * x-ms-original-file: 2025-08-02/Workspaces_Create_MinimumSet_Gen.json
 */
async function workspacesCreateMinimumSetGen() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new IoTFirmwareDefenseClient(credential, subscriptionId);
  const result = await client.workspaces.create("rgworkspaces", "exampleWorkspaceName", {
    properties: {},
    location: "East US",
  });
  console.log(result);
}
