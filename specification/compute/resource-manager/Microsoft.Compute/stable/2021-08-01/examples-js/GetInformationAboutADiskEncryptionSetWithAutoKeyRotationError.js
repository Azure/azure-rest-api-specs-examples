const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function getInformationAboutADiskEncryptionSetWhenAutoKeyRotationFailed() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskEncryptionSetName = "myDiskEncryptionSet";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskEncryptionSets.get(resourceGroupName, diskEncryptionSetName);
  console.log(result);
}

getInformationAboutADiskEncryptionSetWhenAutoKeyRotationFailed().catch(console.error);
