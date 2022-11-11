const { HybridContainerServiceClient } = require("@azure/arm-hybridcontainerservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the hybrid identity metadata proxy resource.
 *
 * @summary Get the hybrid identity metadata proxy resource.
 * x-ms-original-file: specification/hybridaks/resource-manager/Microsoft.HybridContainerService/preview/2022-05-01-preview/examples/GetHybridIdentityMetadata.json
 */
async function getHybridIdentityMetadata() {
  const subscriptionId = "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = "testrg";
  const provisionedClustersName = "ContosoTargetCluster";
  const hybridIdentityMetadataResourceName = "default";
  const credential = new DefaultAzureCredential();
  const client = new HybridContainerServiceClient(credential, subscriptionId);
  const result = await client.hybridIdentityMetadataOperations.get(
    resourceGroupName,
    provisionedClustersName,
    hybridIdentityMetadataResourceName
  );
  console.log(result);
}

getHybridIdentityMetadata().catch(console.error);
