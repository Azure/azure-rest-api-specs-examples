
/**
 * Samples for DiskAccesses GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-02/diskAccessExamples/DiskAccess_Get_WithPrivateEndpoints.json
     */
    /**
     * Sample code: get information about a disk access resource with private endpoints.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getInformationAboutADiskAccessResourceWithPrivateEndpoints(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskAccesses().getByResourceGroupWithResponse("myResourceGroup", "myDiskAccess",
            com.azure.core.util.Context.NONE);
    }
}
