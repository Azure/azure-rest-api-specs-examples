
/**
 * Samples for BackupPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2024-03-01/examples/BackupPolicies_Delete.json
     */
    /**
     * Sample code: BackupPolicies_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void backupPoliciesDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.backupPolicies().delete("resourceGroup", "accountName", "backupPolicyName",
            com.azure.core.util.Context.NONE);
    }
}
