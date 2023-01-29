/** Samples for DiskAccesses GetAPrivateEndpointConnection. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2022-07-02/examples/diskAccessExamples/DiskAccessPrivateEndpointConnection_Get.json
     */
    /**
     * Sample code: Get information about a private endpoint connection under a disk access resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getInformationAboutAPrivateEndpointConnectionUnderADiskAccessResource(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDiskAccesses()
            .getAPrivateEndpointConnectionWithResponse(
                "myResourceGroup", "myDiskAccess", "myPrivateEndpointConnection", com.azure.core.util.Context.NONE);
    }
}
