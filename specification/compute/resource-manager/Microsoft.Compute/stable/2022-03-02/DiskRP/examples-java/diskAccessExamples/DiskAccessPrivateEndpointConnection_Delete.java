import com.azure.core.util.Context;

/** Samples for DiskAccesses DeleteAPrivateEndpointConnection. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-02/DiskRP/examples/diskAccessExamples/DiskAccessPrivateEndpointConnection_Delete.json
     */
    /**
     * Sample code: Delete a private endpoint connection under a disk access resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAPrivateEndpointConnectionUnderADiskAccessResource(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getDiskAccesses()
            .deleteAPrivateEndpointConnection(
                "myResourceGroup", "myDiskAccess", "myPrivateEndpointConnection", Context.NONE);
    }
}
