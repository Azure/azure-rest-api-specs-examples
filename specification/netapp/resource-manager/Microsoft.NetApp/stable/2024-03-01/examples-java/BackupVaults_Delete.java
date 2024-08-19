
/**
 * Samples for BackupVaults Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/examples/BackupVaults_Delete.json
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
