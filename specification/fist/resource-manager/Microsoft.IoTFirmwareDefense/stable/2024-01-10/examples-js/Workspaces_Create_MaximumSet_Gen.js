const { IoTFirmwareDefenseClient } = require("@azure/arm-iotfirmwaredefense");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create or update a firmware analysis workspace.
 *
 * @summary The operation to create or update a firmware analysis workspace.
 * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/Workspaces_Create_MaximumSet_Gen.json
 */
async function workspacesCreateMaximumSetGen() {
  const subscriptionId =
    process.env["IOTFIRMWAREDEFENSE_SUBSCRIPTION_ID"] || "5443A01A-5242-4950-AC1A-2DD362180254";
  const resourceGroupName = process.env["IOTFIRMWAREDEFENSE_RESOURCE_GROUP"] || "rgworkspaces";
  const workspaceName = "E___-3";
  const workspace = {
    location: "jjwbseilitjgdrhbvvkwviqj",
    properties: {},
    tags: { key450: "rzqqumbpfsbibnpirsm" },
  };
  const credential = new DefaultAzureCredential();
  const client = new IoTFirmwareDefenseClient(credential, subscriptionId);
  const result = await client.workspaces.create(resourceGroupName, workspaceName, workspace);
  console.log(result);
}
