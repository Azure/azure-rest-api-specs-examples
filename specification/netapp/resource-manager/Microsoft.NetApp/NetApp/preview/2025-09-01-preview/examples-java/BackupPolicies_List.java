
/**
 * Samples for BackupPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/BackupPolicies_List.json
     */
    /**
     * Sample code: BackupPolicies_List.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupPoliciesList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backupPolicies().list("myRG", "account1", com.azure.core.util.Context.NONE);
    }
}
