
/**
 * Samples for ClusterRecoveryPointOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ClusterRecoveryPoint_Get.json
     */
    /**
     * Sample code: Gets a recovery point.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void
        getsARecoveryPoint(com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.clusterRecoveryPointOperations().getWithResponse("resourceGroupPS1", "vault1", "fabric-pri-eastus",
            "pri-cloud-eastus", "testcluster", "06b9ae7f-f21d-4a76-9897-5cf5d6004d80",
            com.azure.core.util.Context.NONE);
    }
}
