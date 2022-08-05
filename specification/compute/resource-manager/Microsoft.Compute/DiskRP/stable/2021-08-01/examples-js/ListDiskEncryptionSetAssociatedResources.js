const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function listAllResourcesThatAreEncryptedWithThisDiskEncryptionSet() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskEncryptionSetName = "myDiskEncryptionSet";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.diskEncryptionSets.listAssociatedResources(
    resourceGroupName,
    diskEncryptionSetName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllResourcesThatAreEncryptedWithThisDiskEncryptionSet().catch(console.error);
