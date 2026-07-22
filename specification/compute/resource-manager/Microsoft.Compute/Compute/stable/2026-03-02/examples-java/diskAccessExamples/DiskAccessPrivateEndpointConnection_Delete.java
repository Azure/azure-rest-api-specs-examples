
/**
 * Samples for DiskAccesses DeleteAPrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskAccessExamples/DiskAccessPrivateEndpointConnection_Delete.json
     */
    /**
     * Sample code: delete a private endpoint connection under a disk access resource.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void deleteAPrivateEndpointConnectionUnderADiskAccessResource(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskAccesses().deleteAPrivateEndpointConnection("myResourceGroup", "myDiskAccess",
            "myPrivateEndpointConnection", com.azure.core.util.Context.NONE);
    }
}
