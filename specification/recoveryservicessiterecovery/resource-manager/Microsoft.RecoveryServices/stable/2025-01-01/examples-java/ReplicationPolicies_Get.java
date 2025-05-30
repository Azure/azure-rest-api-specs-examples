
/**
 * Samples for ReplicationPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationPolicies_Get.json
     */
    /**
     * Sample code: Gets the requested policy.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        getsTheRequestedPolicy(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationPolicies().getWithResponse("resourceGroupPS1", "vault1", "protectionprofile1",
            com.azure.core.util.Context.NONE);
    }
}
