
/**
 * Samples for ElasticBackupPolicies ListByElasticAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-15-preview/ElasticBackupPolicies_List.json
     */
    /**
     * Sample code: ElasticBackupPolicies_ListByElasticAccount.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void
        elasticBackupPoliciesListByElasticAccount(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticBackupPolicies().listByElasticAccount("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}
