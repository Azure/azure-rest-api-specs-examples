const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function deleteADiskEncryptionSet() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskEncryptionSetName = "myDiskEncryptionSet";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.diskEncryptionSets.beginDeleteAndWait(
    resourceGroupName,
    diskEncryptionSetName
  );
  console.log(result);
}

deleteADiskEncryptionSet().catch(console.error);
