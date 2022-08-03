const { MobileNetworkManagementClient } = require("@azure/arm-mobilenetwork");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a SIM group.
 *
 * @summary Creates or updates a SIM group.
 * x-ms-original-file: specification/mobilenetwork/resource-manager/Microsoft.MobileNetwork/preview/2022-04-01-preview/examples/SimGroupCreate.json
 */
async function createSimGroup() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const simGroupName = "testSimGroup";
  const parameters = {
    encryptionKey: {
      keyUrl: "https://contosovault.vault.azure.net/keys/azureKey",
    },
    location: "eastus",
    mobileNetwork: {
      id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.MobileNetwork/mobileNetworks/testMobileNetwork",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new MobileNetworkManagementClient(credential, subscriptionId);
  const result = await client.simGroups.beginCreateOrUpdateAndWait(
    resourceGroupName,
    simGroupName,
    parameters
  );
  console.log(result);
}

createSimGroup().catch(console.error);
