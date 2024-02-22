
import com.azure.resourcemanager.recoveryservicessiterecovery.models.PossibleOperationsDirections;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.RecoveryPlanHyperVReplicaAzureFailoverInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.RecoveryPlanPlannedFailoverInput;
import com.azure.resourcemanager.recoveryservicessiterecovery.models.RecoveryPlanPlannedFailoverInputProperties;
import java.util.Arrays;

/**
 * Samples for ReplicationRecoveryPlans PlannedFailover.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-08-01/examples
     * /ReplicationRecoveryPlans_PlannedFailover.json
     */
    /**
     * Sample code: Execute planned failover of the recovery plan.
     * 
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void executePlannedFailoverOfTheRecoveryPlan(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager.replicationRecoveryPlans().plannedFailover("vault1", "resourceGroupPS1", "RPtest1",
            new RecoveryPlanPlannedFailoverInput().withProperties(new RecoveryPlanPlannedFailoverInputProperties()
                .withFailoverDirection(PossibleOperationsDirections.PRIMARY_TO_RECOVERY)
                .withProviderSpecificDetails(Arrays.asList(new RecoveryPlanHyperVReplicaAzureFailoverInput()))),
            com.azure.core.util.Context.NONE);
    }
}
