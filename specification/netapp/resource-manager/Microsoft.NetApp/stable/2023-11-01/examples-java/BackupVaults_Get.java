
/**
 * Samples for BackupVaults Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-11-01/examples/BackupVaults_Get.json
     */
    /**
     * Sample code: BackupVaults_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupVaultsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backupVaults().getWithResponse("myRG", "account1", "backupVault1", com.azure.core.util.Context.NONE);
    }
}
