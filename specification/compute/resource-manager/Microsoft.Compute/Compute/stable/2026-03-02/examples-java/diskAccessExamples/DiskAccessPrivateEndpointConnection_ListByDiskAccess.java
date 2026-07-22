
/**
 * Samples for DiskAccesses ListPrivateEndpointConnections.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskAccessExamples/DiskAccessPrivateEndpointConnection_ListByDiskAccess.json
     */
    /**
     * Sample code: get information about a private endpoint connection under a disk access resource.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getInformationAboutAPrivateEndpointConnectionUnderADiskAccessResource(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskAccesses().listPrivateEndpointConnections("myResourceGroup", "myDiskAccess",
            com.azure.core.util.Context.NONE);
    }
}
