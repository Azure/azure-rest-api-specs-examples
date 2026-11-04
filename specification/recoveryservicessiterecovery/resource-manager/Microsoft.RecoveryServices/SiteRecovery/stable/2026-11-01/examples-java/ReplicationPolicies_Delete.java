
/**
 * Samples for ReplicationPolicies Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationPolicies_Delete.json
     */
    /**
     * Sample code: Delete the policy.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        deleteThePolicy(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationPolicies().delete("resourceGroupPS1", "vault1", "protectionprofile1",
            com.azure.core.util.Context.NONE);
    }
}
