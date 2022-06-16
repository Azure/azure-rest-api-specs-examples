const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function getASasOnAManagedDisk() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskName = "myDisk";
  const grantAccessData = {
    access: "Read",
    durationInSeconds: 300,
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

getASasOnAManagedDisk().catch(console.error);
