
import com.azure.resourcemanager.workloadorchestration.models.ActiveState;
import com.azure.resourcemanager.workloadorchestration.models.ExtendedLocation;
import com.azure.resourcemanager.workloadorchestration.models.ExtendedLocationType;
import com.azure.resourcemanager.workloadorchestration.models.InstanceProperties;
import com.azure.resourcemanager.workloadorchestration.models.ReconciliationPolicyProperties;
import com.azure.resourcemanager.workloadorchestration.models.ReconciliationState;

/**
 * Samples for Instances CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Instances_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: Instances_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void instancesCreateOrUpdateMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.instances().define("testname").withExistingSolution("rgconfigurationmanager", "testname", "testname")
            .withProperties(new InstanceProperties().withSolutionVersionId("acpddbkfclsgxg")
                .withTargetId("eguutiftuxrsavvckjrv").withActiveState(ActiveState.ACTIVE)
                .withReconciliationPolicy(new ReconciliationPolicyProperties().withState(ReconciliationState.INACTIVE)
                    .withInterval("szucgzdbydcowvhprhx"))
                .withSolutionScope("testname"))
            .withExtendedLocation(
                new ExtendedLocation().withName("szjrwimeqyiue").withType(ExtendedLocationType.EDGE_ZONE))
            .create();
    }
}
