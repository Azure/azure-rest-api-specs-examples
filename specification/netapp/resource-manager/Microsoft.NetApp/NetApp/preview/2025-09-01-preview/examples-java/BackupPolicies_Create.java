
/**
 * Samples for BackupPolicies Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/BackupPolicies_Create.json
     */
    /**
     * Sample code: BackupPolicies_Create.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupPoliciesCreate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backupPolicies().define("backupPolicyName").withRegion("westus")
            .withExistingNetAppAccount("myRG", "account1").withDailyBackupsToKeep(10).withWeeklyBackupsToKeep(10)
            .withMonthlyBackupsToKeep(10).withEnabled(true).create();
    }
}
