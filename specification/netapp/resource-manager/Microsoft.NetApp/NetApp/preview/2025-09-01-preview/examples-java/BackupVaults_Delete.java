
/**
 * Samples for BackupVaults Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/BackupVaults_Delete.json
     */
    /**
     * Sample code: BackupVaults_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupVaultsDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backupVaults().delete("resourceGroup", "account1", "backupVault1", com.azure.core.util.Context.NONE);
    }
}
