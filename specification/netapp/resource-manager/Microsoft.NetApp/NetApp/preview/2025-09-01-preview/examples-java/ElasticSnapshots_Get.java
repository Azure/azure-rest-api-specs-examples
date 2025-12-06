
/**
 * Samples for ElasticSnapshots Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticSnapshots_Get.json
     */
    /**
     * Sample code: ElasticSnapshots_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticSnapshotsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticSnapshots().getWithResponse("myRG", "account1", "pool1", "volume1", "snapshot1",
            com.azure.core.util.Context.NONE);
    }
}
