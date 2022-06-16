const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function rotateClusterCertificates() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedClusters.beginRotateClusterCertificatesAndWait(
    resourceGroupName,
    resourceName
  );
  console.log(result);
}

rotateClusterCertificates().catch(console.error);
