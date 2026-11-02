
/**
 * Samples for ExternalClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/ExternalClusters_ListBySubscription.json
     */
    /**
     * Sample code: ExternalClusters_ListBySubscription.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void externalClustersListBySubscription(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.externalClusters().list(com.azure.core.util.Context.NONE);
    }
}
