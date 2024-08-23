const { AzureStackHCIClient } = require("@azure/arm-azurestackhci");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Configure RemoteSupport on a cluster
 *
 * @summary Configure RemoteSupport on a cluster
 * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/StackHCI/stable/2024-04-01/examples/ConfigureRemoteSupport.json
 */
async function configureRemoteSupport() {
  const subscriptionId =
    process.env["AZURESTACKHCI_SUBSCRIPTION_ID"] || "fd3c3665-1729-4b7b-9a38-238e83b0f98b";
  const resourceGroupName = process.env["AZURESTACKHCI_RESOURCE_GROUP"] || "test-rg";
  const clusterName = "mycluster";
  const remoteSupportRequest = {
    properties: {
      accessLevel: "Diagnostics",
      expirationTimeStamp: new Date("2020-01-01T17:18:19.1234567Z"),
      remoteSupportType: "Enable",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureStackHCIClient(credential, subscriptionId);
  const result = await client.clusters.beginConfigureRemoteSupportAndWait(
    resourceGroupName,
    clusterName,
    remoteSupportRequest,
  );
  console.log(result);
}
