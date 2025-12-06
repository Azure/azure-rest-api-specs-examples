
/**
 * Samples for ElasticBackupVaults CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticBackupVaults_Create.json
     */
    /**
     * Sample code: ElasticBackupVaults_CreateOrUpdate.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticBackupVaultsCreateOrUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticBackupVaults().define("backupVault1").withRegion("eastus")
            .withExistingElasticAccount("myRG", "account1").create();
    }
}
