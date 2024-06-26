const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Collect a diagnostics package for the specified packet core control plane. This action will upload the diagnostics to a storage account.
 *
 * @summary Collect a diagnostics package for the specified packet core control plane. This action will upload the diagnostics to a storage account.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2024-04-01/examples/PacketCoreControlPlaneCollectDiagnosticsPackage.json
 */
async function collectDiagnosticsPackageFromPacketCoreControlPlane() {
  const subscriptionId =
    process.env["MOBILENETWORK_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MOBILENETWORK_RESOURCE_GROUP"] || "rg1";
  const packetCoreControlPlaneName = "TestPacketCoreCP";
  const parameters = {
    storageAccountBlobUrl:
      "https://contosoaccount.blob.core.windows.net/container/diagnosticsPackage.zip",
  };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.packetCoreControlPlanes.beginCollectDiagnosticsPackageAndWait(
    resourceGroupName,
    packetCoreControlPlaneName,
    parameters,
  );
  console.log(result);
}
