
/**
 * Samples for Volumes SplitCloneFromParent.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-07-01-preview/examples/Volumes_SplitClone.
     * json
     */
    /**
     * Sample code: Volumes_SplitClone.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesSplitClone(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().splitCloneFromParent("myRG", "account1", "pool1", "volume1",
            com.azure.core.util.Context.NONE);
    }
}
