const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function tableServicesGet() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4410";
  const accountName = "sto8607";
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.tableServices.getServiceProperties(resourceGroupName, accountName);
  console.log(result);
}

tableServicesGet().catch(console.error);
