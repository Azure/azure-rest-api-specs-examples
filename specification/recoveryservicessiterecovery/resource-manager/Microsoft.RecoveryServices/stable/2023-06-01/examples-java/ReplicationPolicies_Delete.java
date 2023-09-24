/** Samples for ReplicationPolicies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationPolicies_Delete.json
     */
    /**
     * Sample code: Delete the policy.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void deleteThePolicy(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .replicationPolicies()
            .delete("vault1", "resourceGroupPS1", "protectionprofile1", com.azure.core.util.Context.NONE);
    }
}
