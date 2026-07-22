
/**
 * Samples for DiskAccesses GetPrivateLinkResources.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskAccessExamples/DiskAccessPrivateLinkResources_Get.json
     */
    /**
     * Sample code: list all possible private link resources under disk access resource.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listAllPossiblePrivateLinkResourcesUnderDiskAccessResource(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskAccesses().getPrivateLinkResourcesWithResponse("myResourceGroup", "myDiskAccess",
            com.azure.core.util.Context.NONE);
    }
}
