
/**
 * Samples for DiskAccesses Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskAccessExamples/DiskAccess_Delete.json
     */
    /**
     * Sample code: delete a disk access resource.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void deleteADiskAccessResource(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskAccesses().delete("myResourceGroup", "myDiskAccess",
            com.azure.core.util.Context.NONE);
    }
}
