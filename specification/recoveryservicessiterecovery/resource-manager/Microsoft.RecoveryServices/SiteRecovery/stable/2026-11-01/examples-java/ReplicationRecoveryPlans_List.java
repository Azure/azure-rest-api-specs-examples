
/**
 * Samples for ReplicationRecoveryPlans List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ReplicationRecoveryPlans_List.json
     */
    /**
     * Sample code: Gets the list of recovery plans.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        getsTheListOfRecoveryPlans(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationRecoveryPlans().list("resourceGroupPS1", "vault1", com.azure.core.util.Context.NONE);
    }
}
