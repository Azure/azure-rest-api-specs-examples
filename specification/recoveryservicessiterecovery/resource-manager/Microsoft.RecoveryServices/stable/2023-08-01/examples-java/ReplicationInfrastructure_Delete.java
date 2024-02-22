
/**
 * Samples for ReplicationFabrics RemoveInfra.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationInfrastructure_Delete.json
     */
    /**
     * Sample code: Removes the appliance's infrastructure under the fabric.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void removesTheApplianceSInfrastructureUnderTheFabric(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationFabrics().removeInfra("resourceGroupPS1", "vault1", "cloud1",
            com.azure.core.util.Context.NONE);
    }
}
