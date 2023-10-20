/** Samples for BackupPolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-05-01/examples/BackupPolicies_Delete.json
     */
    /**
     * Sample code: Backups_Delete.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupsDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager
            .backupPolicies()
            .delete("resourceGroup", "accountName", "backupPolicyName", com.azure.core.util.Context.NONE);
    }
}
