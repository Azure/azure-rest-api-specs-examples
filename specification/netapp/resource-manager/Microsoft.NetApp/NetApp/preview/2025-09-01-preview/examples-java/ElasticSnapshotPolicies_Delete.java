
/**
 * Samples for ElasticSnapshotPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/ElasticSnapshotPolicies_Delete.json
     */
    /**
     * Sample code: ElasticSnapshotPolicies_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void elasticSnapshotPoliciesDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.elasticSnapshotPolicies().delete("resourceGroup", "accountName", "snapshotPolicyName",
            com.azure.core.util.Context.NONE);
    }
}
