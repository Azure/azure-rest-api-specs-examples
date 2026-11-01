
/**
 * Samples for ExternalCacheConnections List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/ExternalCacheConnections_ListBySubscription.json
     */
    /**
     * Sample code: ExternalCacheConnections_ListBySubscription.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void
        externalCacheConnectionsListBySubscription(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.externalCacheConnections().list(com.azure.core.util.Context.NONE);
    }
}
