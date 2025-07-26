
/**
 * Samples for DiskAccesses GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2025-01-02/examples/diskAccessExamples/
     * DiskAccess_Get.json
     */
    /**
     * Sample code: get information about a disk access resource.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getInformationAboutADiskAccessResource(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDiskAccesses()
            .getByResourceGroupWithResponse("myResourceGroup", "myDiskAccess", com.azure.core.util.Context.NONE);
    }
}
