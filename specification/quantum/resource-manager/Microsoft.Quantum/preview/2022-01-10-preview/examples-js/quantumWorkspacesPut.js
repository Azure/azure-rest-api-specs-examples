const { AzureQuantumManagementClient } = require("@azure/arm-quantum");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a workspace resource.
 *
 * @summary Creates or updates a workspace resource.
 * x-ms-original-file: specification/quantum/resource-manager/Microsoft.Quantum/preview/2022-01-10-preview/examples/quantumWorkspacesPut.json
 */
async function quantumWorkspacesPut() {
  const subscriptionId =
    process.env["QUANTUM_SUBSCRIPTION_ID"] || "00000000-1111-2222-3333-444444444444";
  const resourceGroupName = process.env["QUANTUM_RESOURCE_GROUP"] || "quantumResourcegroup";
  const workspaceName = "quantumworkspace1";
  const quantumWorkspace = {
    location: "West US",
    providers: [
      { providerId: "Honeywell", providerSku: "Basic" },
      { providerId: "IonQ", providerSku: "Basic" },
      { providerId: "OneQBit", providerSku: "Basic" },
    ],
    storageAccount:
      "/subscriptions/1C4B2828-7D49-494F-933D-061373BE28C2/resourceGroups/quantumResourcegroup/providers/Microsoft.Storage/storageAccounts/testStorageAccount",
  };
  const credential = new DefaultAzureCredential();
  const client = new AzureQuantumManagementClient(credential, subscriptionId);
  const result = await client.workspaces.beginCreateOrUpdateAndWait(
    resourceGroupName,
    workspaceName,
    quantumWorkspace
  );
  console.log(result);
}
