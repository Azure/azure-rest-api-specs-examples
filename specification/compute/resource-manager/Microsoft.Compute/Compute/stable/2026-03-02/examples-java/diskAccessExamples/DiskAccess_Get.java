
/**
 * Samples for DiskAccesses GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskAccessExamples/DiskAccess_Get.json
     */
    /**
     * Sample code: get information about a disk access resource.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        getInformationAboutADiskAccessResource(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskAccesses().getByResourceGroupWithResponse("myResourceGroup", "myDiskAccess",
            com.azure.core.util.Context.NONE);
    }
}
