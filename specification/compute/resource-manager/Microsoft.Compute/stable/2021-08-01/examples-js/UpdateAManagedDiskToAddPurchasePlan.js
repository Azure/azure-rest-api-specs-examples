const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function updateAManagedDiskToAddPurchasePlan() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const diskName = "myDisk";
  const disk = {
    purchasePlan: {
      name: "myPurchasePlanName",
      product: "myPurchasePlanProduct",
      promotionCode: "myPurchasePlanPromotionCode",
      publisher: "myPurchasePlanPublisher",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.beginUpdateAndWait(resourceGroupName, diskName, disk);
  console.log(result);
}

updateAManagedDiskToAddPurchasePlan().catch(console.error);
