const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

async function createOrUpdateManagedNamespace() {
  const subscriptionId =
    process.env["CONTAINERSERVICE_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["CONTAINERSERVICE_RESOURCE_GROUP"] || "rg1";
  const resourceName = "clustername1";
  const managedNamespaceName = "namespace1";
  const parameters = {
    properties: {
      adoptionPolicy: "IfIdentical",
      annotations: { annatationKey: "annatationValue" },
      defaultNetworkPolicy: {
        egress: "AllowAll",
        ingress: "AllowSameNamespace",
      },
      defaultResourceQuota: {
        cpuLimit: "3m",
        cpuRequest: "3m",
        memoryLimit: "5Gi",
        memoryRequest: "5Gi",
      },
      deletePolicy: "Keep",
      labels: { "kubernetesIo/metadataName": "true" },
    },
    tags: { tagKey1: "tagValue1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.managedNamespaces.beginCreateOrUpdateAndWait(
    resourceGroupName,
    resourceName,
    managedNamespaceName,
    parameters,
  );
  console.log(result);
}
