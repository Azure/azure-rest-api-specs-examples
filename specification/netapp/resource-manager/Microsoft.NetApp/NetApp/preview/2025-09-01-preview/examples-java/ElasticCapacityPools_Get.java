
/**
 * Samples for ElasticCapacityPools Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticCapacityPools_Get.json
     */
    /**
     * Sample code: ElasticCapacityPools_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticCapacityPoolsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticCapacityPools().getWithResponse("myRG", "account1", "pool1", com.azure.core.util.Context.NONE);
    }
}
