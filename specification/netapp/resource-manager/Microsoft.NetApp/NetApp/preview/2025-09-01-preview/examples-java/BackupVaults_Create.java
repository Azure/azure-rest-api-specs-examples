
/**
 * Samples for BackupVaults CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/BackupVaults_Create.json
     */
    /**
     * Sample code: BackupVault_CreateOrUpdate.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupVaultCreateOrUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backupVaults().define("backupVault1").withRegion("eastus").withExistingNetAppAccount("myRG", "account1")
            .create();
    }
}
