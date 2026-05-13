
/**
 * Samples for DeletedBackupVaults ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/DeletedBackupVaults_ListByLocation.json
     */
    /**
     * Sample code: List deleted backup vaults by location.
     * 
     * @param manager Entry point to DataProtectionManager.
     */
    public static void
        listDeletedBackupVaultsByLocation(com.azure.resourcemanager.dataprotection.DataProtectionManager manager) {
        manager.deletedBackupVaults().listByLocation("westus", com.azure.core.util.Context.NONE);
    }
}
