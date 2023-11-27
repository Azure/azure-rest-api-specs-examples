const { HybridComputeManagementClient } = require("@azure/arm-hybridcompute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Implements HybridIdentityMetadata GET method.
 *
 * @summary Implements HybridIdentityMetadata GET method.
 * x-ms-original-file: specification/hybridcompute/resource-manager/Microsoft.HybridCompute/preview/2023-06-20-preview/examples/HybridIdentityMetadata_Get.json
 */
async function getHybridIdentityMetadata() {
  const subscriptionId =
    process.env["HYBRIDCOMPUTE_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["HYBRIDCOMPUTE_RESOURCE_GROUP"] || "testrg";
  const machineName = "ContosoVm";
  const metadataName = "default";
  const credential = new DefaultAzureCredential();
  const client = new HybridComputeManagementClient(credential, subscriptionId);
  const result = await client.hybridIdentityMetadataOperations.get(
    resourceGroupName,
    machineName,
    metadataName
  );
  console.log(result);
}
