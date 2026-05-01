const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to updates (patches) a disk.
 *
 * @summary updates (patches) a disk.
 * x-ms-original-file: 2025-01-02/diskExamples/Disk_Update_AddPurchasePlan.json
 */
async function updateAManagedDiskToAddPurchasePlan() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.disks.update("myResourceGroup", "myDisk", {
    purchasePlan: {
      name: "myPurchasePlanName",
      publisher: "myPurchasePlanPublisher",
      product: "myPurchasePlanProduct",
      promotionCode: "myPurchasePlanPromotionCode",
    },
  });
  console.log(result);
}
