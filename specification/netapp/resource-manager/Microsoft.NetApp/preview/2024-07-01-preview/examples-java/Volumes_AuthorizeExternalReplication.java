
/**
 * Samples for Volumes AuthorizeExternalReplication.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2024-07-01-preview/examples/
     * Volumes_AuthorizeExternalReplication.json
     */
    /**
     * Sample code: Volumes_AuthorizeExternalReplication.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void
        volumesAuthorizeExternalReplication(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().authorizeExternalReplication("myRG", "account1", "pool1", "volume1",
            com.azure.core.util.Context.NONE);
    }
}
