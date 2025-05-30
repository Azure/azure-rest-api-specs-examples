
import com.azure.resourcemanager.recoveryservicessiterecovery.models.PossibleOperationsDirections;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.RecoveryPlanHyperVReplicaAzureFailoverInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.RecoveryPlanTestFailoverInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.RecoveryPlanTestFailoverInputProperties;
import java.util.Arrays;

/**
 * Samples for ReplicationRecoveryPlans TestFailover.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2025-01-01/examples
     * /ReplicationRecoveryPlans_TestFailover.json
     */
    /**
     * Sample code: Execute test failover of the recovery plan.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void executeTestFailoverOfTheRecoveryPlan(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationRecoveryPlans().testFailover("resourceGroupPS1", "vault1", "RPtest1",
            new RecoveryPlanTestFailoverInput().withProperties(new RecoveryPlanTestFailoverInputProperties()
                .withFailoverDirection(PossibleOperationsDirections.PRIMARY_TO_RECOVERY)
                .withNetworkType("VmNetworkAsInput")
                .withNetworkId(
                    "/subscriptions/c183865e-6077-46f2-a3b1-deb0f4f4650a/resourceGroups/siterecoveryProd1/providers/Microsoft.Network/virtualNetworks/vnetavrai")
                .withProviderSpecificDetails(Arrays.asList(new RecoveryPlanHyperVReplicaAzureFailoverInput()))),
            com.azure.core.util.Context.NONE);
    }
}
