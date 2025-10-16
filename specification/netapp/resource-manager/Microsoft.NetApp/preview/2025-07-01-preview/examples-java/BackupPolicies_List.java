
/**
 * Samples for BackupPolicies List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-07-01-preview/examples/BackupPolicies_List.
     * json
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
