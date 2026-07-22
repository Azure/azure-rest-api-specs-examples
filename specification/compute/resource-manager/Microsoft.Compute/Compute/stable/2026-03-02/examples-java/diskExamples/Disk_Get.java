
/**
 * Samples for Disks GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskExamples/Disk_Get.json
     */
    /**
     * Sample code: get information about a managed disk.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getInformationAboutAManagedDisk(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().getByResourceGroupWithResponse("myResourceGroup", "myManagedDisk",
            com.azure.core.util.Context.NONE);
    }
}
