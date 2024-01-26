
/**
 * Samples for DiskAccesses GetPrivateLinkResources.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-04-02/examples/diskAccessExamples/
     * DiskAccessPrivateLinkResources_Get.json
     */
    /**
     * Sample code: List all possible private link resources under disk access resource.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listAllPossiblePrivateLinkResourcesUnderDiskAccessResource(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDiskAccesses()
            .getPrivateLinkResourcesWithResponse("myResourceGroup", "myDiskAccess", com.azure.core.util.Context.NONE);
    }
}
