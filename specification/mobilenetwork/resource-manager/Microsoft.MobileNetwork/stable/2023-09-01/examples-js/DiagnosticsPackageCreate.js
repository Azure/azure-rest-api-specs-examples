const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a diagnostics package.
 *
 * @summary Creates or updates a diagnostics package.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-09-01/examples/DiagnosticsPackageCreate.json
 */
async function createDiagnosticsPackage() {
  const subscriptionId =
    process.env["MOBILENETWORK_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MOBILENETWORK_RESOURCE_GROUP"] || "rg1";
  const packetCoreControlPlaneName = "TestPacketCoreCP";
  const diagnosticsPackageName = "dp1";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.diagnosticsPackages.beginCreateOrUpdateAndWait(
    resourceGroupName,
    packetCoreControlPlaneName,
    diagnosticsPackageName
  );
  console.log(result);
}
