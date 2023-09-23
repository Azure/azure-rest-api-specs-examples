/** Samples for ReplicationRecoveryPlans List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/ReplicationRecoveryPlans_List.json
     */
    /**
     * Sample code: Gets the list of recovery plans.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheListOfRecoveryPlans(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationRecoveryPlans().list("vault1", "resourceGroupPS1", com.azure.core.util.Context.NONE);
    }
}
