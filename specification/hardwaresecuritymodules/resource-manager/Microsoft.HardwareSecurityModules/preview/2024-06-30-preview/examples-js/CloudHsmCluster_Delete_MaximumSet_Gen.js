const { AzureHSMResourceProvider } = require("@azure/arm-hardwaresecuritymodules");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified Cloud HSM Cluster
 *
 * @summary Deletes the specified Cloud HSM Cluster
 * x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-preview/examples/CloudHsmCluster_Delete_MaximumSet_Gen.json
 */
async function cloudHsmClusterDeleteMaximumSetGen() {
  const subscriptionId =
    process.env["HARDWARESECURITYMODULES_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["HARDWARESECURITYMODULES_RESOURCE_GROUP"] || "rgcloudhsm";
  const cloudHsmClusterName = "chsm1";
  const credential = new DefaultAzureCredential();
  const client = new AzureHSMResourceProvider(credential, subscriptionId);
  const result = await client.cloudHsmClusters.beginDeleteAndWait(
    resourceGroupName,
    cloudHsmClusterName,
  );
  console.log(result);
}
