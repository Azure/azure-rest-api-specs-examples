
/**
 * Samples for ExternalCacheConnections Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/ExternalCacheConnections_Delete.json
     */
    /**
     * Sample code: ExternalCacheConnections_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void externalCacheConnectionsDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.externalCacheConnections().delete("contosoNetAppRg", "contosoCacheConnection",
            com.azure.core.util.Context.NONE);
    }
}
