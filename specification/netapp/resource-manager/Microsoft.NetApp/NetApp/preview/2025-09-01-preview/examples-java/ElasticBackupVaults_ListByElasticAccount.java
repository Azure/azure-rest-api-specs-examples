
/**
 * Samples for ElasticBackupVaults ListByElasticAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticBackupVaults_ListByElasticAccount.json
     */
    /**
     * Sample code: ElasticBackupVaults_ListByElasticAccount.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void
        elasticBackupVaultsListByElasticAccount(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticBackupVaults().listByElasticAccount("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}
