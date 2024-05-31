
/**
 * Samples for BackupVaults CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-11-01/examples/BackupVaults_Create.json
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
