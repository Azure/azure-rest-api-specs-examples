
/**
 * Samples for ElasticSnapshots Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticSnapshots_Delete.json
     */
    /**
     * Sample code: ElasticSnapshots_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticSnapshotsDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticSnapshots().delete("myRG", "account1", "pool1", "volume1", "snapshot1",
            com.azure.core.util.Context.NONE);
    }
}
