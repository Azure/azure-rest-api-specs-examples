
/**
 * Samples for ElasticCapacityPools Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticCapacityPools_Delete.json
     */
    /**
     * Sample code: ElasticCapacityPools_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticCapacityPoolsDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticCapacityPools().delete("myRG", "account1", "pool1", com.azure.core.util.Context.NONE);
    }
}
