const { IoTFirmwareDefenseClient } = require("@azure/arm-iotfirmwaredefense");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a firmware analysis workspace.
 *
 * @summary the operation to create or update a firmware analysis workspace.
 * x-ms-original-file: 2025-08-02/Workspaces_Create_MaximumSet_Gen.json
 */
async function workspacesCreateMaximumSetGenGeneratedByMaximumSetRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new IoTFirmwareDefenseClient(credential, subscriptionId);
  const result = await client.workspaces.create("rgiotfirmwaredefense", "exampleWorkspaceName", {
    properties: {},
    tags: { key4630: "rov" },
    location: "East US",
    sku: {
      name: "Free",
      tier: "Free",
      size: "Free",
      family: "F",
      capacity: 30,
    },
  });
  console.log(result);
}
