import com.azure.resourcemanager.recoveryservicesdatareplication.fluent.models.PlannedFailoverModelInner;
import com.azure.resourcemanager.recoveryservicesdatareplication.models.PlannedFailoverModelCustomProperties;
import com.azure.resourcemanager.recoveryservicesdatareplication.models.PlannedFailoverModelProperties;

/** Samples for ProtectedItem PlannedFailover. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicesdatareplication/resource-manager/Microsoft.DataReplication/preview/2021-02-16-preview/examples/ProtectedItem_PlannedFailover.json
     */
    /**
     * Sample code: ProtectedItem_PlannedFailover.
     *
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void protectedItemPlannedFailover(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager
            .protectedItems()
            .plannedFailover(
                "rgrecoveryservicesdatareplication",
                "4",
                "d",
                new PlannedFailoverModelInner()
                    .withProperties(
                        new PlannedFailoverModelProperties()
                            .withCustomProperties(new PlannedFailoverModelCustomProperties())),
                com.azure.core.util.Context.NONE);
    }
}
