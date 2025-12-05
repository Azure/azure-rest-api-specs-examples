
/**
 * Samples for ElasticBackupPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticBackupPolicies_Delete.json
     */
    /**
     * Sample code: ElasticBackupPolicies_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticBackupPoliciesDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticBackupPolicies().delete("resourceGroup", "accountName", "backupPolicyName",
            com.azure.core.util.Context.NONE);
    }
}
