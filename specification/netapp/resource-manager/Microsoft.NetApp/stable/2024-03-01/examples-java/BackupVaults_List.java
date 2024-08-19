
/**
 * Samples for BackupVaults ListByNetAppAccount.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/examples/BackupVaults_List.json
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
