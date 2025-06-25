const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to The operation to update a license profile.
 *
 * @summary The operation to update a license profile.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2025-02-19-preview/examples/licenseProfile/LicenseProfile_Update.json
 */
async function updateALicenseProfile() {
  const subscriptionId = process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "{subscriptionId}";
  const resourceGroupName = process.env["HYBRIDCOMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const machineName = "myMachine";
  const parameters = {
    assignedLicense: "{LicenseResourceId}",
    productFeatures: [{ name: "Hotpatch", subscriptionStatus: "Enable" }],
    productType: "WindowsServer",
    softwareAssuranceCustomer: true,
    subscriptionStatus: "Enable",
  };
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.licenseProfiles.beginUpdateAndWait(
    resourceGroupName,
    machineName,
    parameters,
  );
  console.log(result);
}
