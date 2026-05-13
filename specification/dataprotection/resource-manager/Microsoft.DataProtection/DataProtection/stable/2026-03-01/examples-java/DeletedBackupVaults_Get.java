
/**
 * Samples for DeletedBackupVaults Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/DeletedBackupVaults_Get.json
     */
    /**
     * Sample code: Get a deleted backup vault.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void getADeletedBackupVault(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.deletedBackupVaults().getWithResponse("westus", "deleted-vault-01", com.azure.core.util.Context.NONE);
    }
}
