const { DataBoxEdgeManagementClient } = require("@azure/arm-databoxedge");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new StorageAccount or updates an existing StorageAccount on the device.
 *
 * @summary Creates a new StorageAccount or updates an existing StorageAccount on the device.
 * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2021-06-01/examples/StorageAccountPut.json
 */
async function storageAccountPut() {
  const subscriptionId = "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const deviceName = "testedgedevice";
  const storageAccountName = "blobstorageaccount1";
  const resourceGroupName = "GroupForEdgeAutomation";
  const storageAccount = {
    description: "It's an awesome storage account",
    dataPolicy: "Cloud",
    storageAccountCredentialId:
      "/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForDataBoxEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/storageAccountCredentials/cisbvt",
    storageAccountStatus: "OK",
  };
  const credential = new DefaultAzureCredential();
  const client = new DataBoxEdgeManagementClient(credential, subscriptionId);
  const result = await client.storageAccounts.beginCreateOrUpdateAndWait(
    deviceName,
    storageAccountName,
    resourceGroupName,
    storageAccount
  );
  console.log(result);
}

storageAccountPut().catch(console.error);
