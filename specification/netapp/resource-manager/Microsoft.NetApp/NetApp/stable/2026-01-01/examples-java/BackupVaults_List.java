
/**
 * Samples for BackupVaults ListByNetAppAccount.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/BackupVaults_List.json
     */
    /**
     * Sample code: BackupVaults_List.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupVaultsList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backupVaults().listByNetAppAccount("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}
