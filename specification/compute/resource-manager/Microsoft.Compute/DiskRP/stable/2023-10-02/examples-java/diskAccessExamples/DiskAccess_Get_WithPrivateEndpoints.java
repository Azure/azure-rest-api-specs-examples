
/**
 * Samples for DiskAccesses GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2023-10-02/examples/diskAccessExamples/
     * DiskAccess_Get_WithPrivateEndpoints.json
     */
    /**
     * Sample code: Get information about a disk access resource with private endpoints.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getInformationAboutADiskAccessResourceWithPrivateEndpoints(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDiskAccesses()
            .getByResourceGroupWithResponse("myResourceGroup", "myDiskAccess", com.azure.core.util.Context.NONE);
    }
}
