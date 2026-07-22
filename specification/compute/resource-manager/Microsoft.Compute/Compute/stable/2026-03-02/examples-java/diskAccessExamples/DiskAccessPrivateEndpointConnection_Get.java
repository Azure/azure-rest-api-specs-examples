
/**
 * Samples for DiskAccesses GetAPrivateEndpointConnection.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-02/diskAccessExamples/DiskAccessPrivateEndpointConnection_Get.json
     */
    /**
     * Sample code: get information about a private endpoint connection under a disk access resource.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void getInformationAboutAPrivateEndpointConnectionUnderADiskAccessResource(
        com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getDiskAccesses().getAPrivateEndpointConnectionWithResponse("myResourceGroup",
            "myDiskAccess", "myPrivateEndpointConnection", com.azure.core.util.Context.NONE);
    }
}
