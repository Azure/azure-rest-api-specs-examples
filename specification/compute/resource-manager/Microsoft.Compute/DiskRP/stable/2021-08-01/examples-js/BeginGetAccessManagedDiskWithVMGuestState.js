const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function getSasOnManagedDiskAndVMGuestState() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskName = "myDisk";
  const grantAccessData = {
    access: "Read",
    durationInSeconds: 300,
    getSecureVMGuestStateSAS: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.beginGrantAccessAndWait(
    resourceGroupName,
    diskName,
    grantAccessData
  );
  console.log(result);
}

getSasOnManagedDiskAndVMGuestState().catch(console.error);
