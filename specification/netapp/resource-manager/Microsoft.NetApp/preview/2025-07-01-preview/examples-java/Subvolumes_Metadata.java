
/**
 * Samples for Subvolumes GetMetadata.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-07-01-preview/examples/Subvolumes_Metadata.
     * json
     */
    /**
     * Sample code: Subvolumes_Metadata.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void subvolumesMetadata(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.subvolumes().getMetadata("myRG", "account1", "pool1", "volume1", "subvolume1",
            com.azure.core.util.Context.NONE);
    }
}
