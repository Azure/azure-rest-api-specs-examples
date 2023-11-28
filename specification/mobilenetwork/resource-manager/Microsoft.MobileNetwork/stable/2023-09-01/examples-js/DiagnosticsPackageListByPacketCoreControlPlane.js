const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all the diagnostics packages under a packet core control plane.
 *
 * @summary Lists all the diagnostics packages under a packet core control plane.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/stable/2023-09-01/examples/DiagnosticsPackageListByPacketCoreControlPlane.json
 */
async function listDiagnosticsPackagesUnderAPacketCoreControlPlane() {
  const subscriptionId =
    process.env["MOBILENETWORK_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["MOBILENETWORK_RESOURCE_GROUP"] || "rg1";
  const packetCoreControlPlaneName = "TestPacketCoreCP";
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.diagnosticsPackages.listByPacketCoreControlPlane(
    resourceGroupName,
    packetCoreControlPlaneName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}
