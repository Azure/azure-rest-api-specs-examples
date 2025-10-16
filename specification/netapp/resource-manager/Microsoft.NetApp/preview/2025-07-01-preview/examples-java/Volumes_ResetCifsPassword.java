
/**
 * Samples for Volumes ResetCifsPassword.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-07-01-preview/examples/
     * Volumes_ResetCifsPassword.json
     */
    /**
     * Sample code: Volumes_ResetCifsPassword.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumesResetCifsPassword(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().resetCifsPassword("myRG", "account1", "pool1", "volume1", com.azure.core.util.Context.NONE);
    }
}
