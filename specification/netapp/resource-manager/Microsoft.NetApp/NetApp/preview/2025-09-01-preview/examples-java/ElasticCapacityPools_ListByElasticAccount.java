
/**
 * Samples for ElasticCapacityPools ListByElasticAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticCapacityPools_ListByElasticAccount.json
     */
    /**
     * Sample code: ElasticCapacityPools_ListByElasticAccount.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void
        elasticCapacityPoolsListByElasticAccount(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticCapacityPools().listByElasticAccount("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}
