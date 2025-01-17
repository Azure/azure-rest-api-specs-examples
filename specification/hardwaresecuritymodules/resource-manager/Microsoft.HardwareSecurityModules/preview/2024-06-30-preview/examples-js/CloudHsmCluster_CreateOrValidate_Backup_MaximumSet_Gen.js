const { AzureHSMResourceProvider } = require("@azure/arm-hardwaresecuritymodules");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Pre Backup operation to validate whether the customer can perform a backup on the Cloud HSM Cluster resource in the specified subscription.
 *
 * @summary Pre Backup operation to validate whether the customer can perform a backup on the Cloud HSM Cluster resource in the specified subscription.
 * x-ms-original-file: specification/hardwaresecuritymodules/resource-manager/Microsoft.HardwareSecurityModules/preview/2024-06-30-preview/examples/CloudHsmCluster_CreateOrValidate_Backup_MaximumSet_Gen.json
 */
async function cloudHsmClusterValidateBackupValidationMaximumSetGen() {
  const subscriptionId =
    process.env["HARDWARESECURITYMODULES_SUBSCRIPTION_ID"] ||
    "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["HARDWARESECURITYMODULES_RESOURCE_GROUP"] || "rgcloudhsm";
  const cloudHsmClusterName = "chsm1";
  const credential = new DefaultAzureCredential();
  const client = new AzureHSMResourceProvider(credential, subscriptionId);
  const result = await client.cloudHsmClusters.beginValidateBackupPropertiesAndWait(
    resourceGroupName,
    cloudHsmClusterName,
  );
  console.log(result);
}
