
/**
 * Samples for Disks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskExamples/Disk_Delete.json
     */
    /**
     * Sample code: delete a managed disk.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void deleteAManagedDisk(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDisks().delete("myResourceGroup", "myDisk", com.azure.core.util.Context.NONE);
    }
}
