
/**
 * Samples for ElasticVolumes ListByElasticPool.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticVolumes_ListByElasticPool.json
     */
    /**
     * Sample code: ElasticVolumes_ListByElasticPool.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticVolumesListByElasticPool(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticVolumes().listByElasticPool("myRG", "account1", "pool1", com.azure.core.util.Context.NONE);
    }
}
