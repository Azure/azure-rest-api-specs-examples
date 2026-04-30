
import com.azure.resourcemanager.compute.models.DiskPurchasePlan;
import com.azure.resourcemanager.compute.models.DiskUpdate;

/**
 * Samples for Disks Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskExamples/Disk_Update_AddPurchasePlan.json
     */
    /**
     * Sample code: update a managed disk to add purchase plan.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void updateAManagedDiskToAddPurchasePlan(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().update("myResourceGroup", "myDisk",
            new DiskUpdate().withPurchasePlan(
                new DiskPurchasePlan().withName("myPurchasePlanName").withPublisher("myPurchasePlanPublisher")
                    .withProduct("myPurchasePlanProduct").withPromotionCode("fakeTokenPlaceholder")),
            com.azure.core.util.Context.NONE);
    }
}
