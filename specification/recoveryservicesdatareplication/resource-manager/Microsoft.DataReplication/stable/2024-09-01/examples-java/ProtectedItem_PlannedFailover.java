
import com.azure.resourcemanager.recoveryservicesdatareplication.fluent.models.PlannedFailoverModelInner;
import com.azure.resourcemanager.recoveryservicesdatareplication.models.PlannedFailoverModelCustomProperties;
import com.azure.resourcemanager.recoveryservicesdatareplication.models.PlannedFailoverModelProperties;

/**
 * Samples for ProtectedItem PlannedFailover.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/ProtectedItem_PlannedFailover.json
     */
    /**
     * Sample code: Performs planned failover.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void performsPlannedFailover(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.protectedItems().plannedFailover("rgrecoveryservicesdatareplication", "4", "d",
            new PlannedFailoverModelInner().withProperties(
                new PlannedFailoverModelProperties().withCustomProperties(new PlannedFailoverModelCustomProperties())),
            com.azure.core.util.Context.NONE);
    }
}
