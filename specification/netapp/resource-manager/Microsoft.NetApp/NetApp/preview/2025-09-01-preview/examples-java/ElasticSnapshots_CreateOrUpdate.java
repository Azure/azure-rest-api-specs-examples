
/**
 * Samples for ElasticSnapshots CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticSnapshots_CreateOrUpdate.json
     */
    /**
     * Sample code: ElasticSnapshots_CreateOrUpdate.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticSnapshotsCreateOrUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticSnapshots().define("snapshot1").withExistingElasticVolume("myRG", "account1", "pool1", "volume1")
            .create();
    }
}
