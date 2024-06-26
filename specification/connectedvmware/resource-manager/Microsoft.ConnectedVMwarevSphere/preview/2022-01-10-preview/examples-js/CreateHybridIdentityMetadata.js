const { AzureArcVMwareManagementServiceAPI } = require("@azure/arm-connectedvmware");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create Or Update HybridIdentityMetadata.
 *
 * @summary Create Or Update HybridIdentityMetadata.
 * x-ms-original-file: specification/connectedvmware/resource-manager/Microsoft.ConnectedVMwarevSphere/preview/2022-01-10-preview/examples/CreateHybridIdentityMetadata.json
 */
async function createHybridIdentityMetadata() {
  const subscriptionId =
    process.env["CONNECTEDVMWARE_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["CONNECTEDVMWARE_RESOURCE_GROUP"] || "testrg";
  const virtualMachineName = "ContosoVm";
  const metadataName = "default";
  const body = {
    publicKey: "8ec7d60c-9700-40b1-8e6e-e5b2f6f477f2",
    vmId: "f8b82dff-38ef-4220-99ef-d3a3f86ddc6c",
  };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new AzureArcVMwareManagementServiceAPI(credential, subscriptionId);
  const result = await client.hybridIdentityMetadataOperations.create(
    resourceGroupName,
    virtualMachineName,
    metadataName,
    options
  );
  console.log(result);
}
