
import com.azure.resourcemanager.workloadorchestration.models.ActiveState;
import com.azure.resourcemanager.workloadorchestration.models.Instance;
import com.azure.resourcemanager.workloadorchestration.models.InstanceProperties;
import com.azure.resourcemanager.workloadorchestration.models.ReconciliationPolicyProperties;
import com.azure.resourcemanager.workloadorchestration.models.ReconciliationState;

/**
 * Samples for Instances Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Instances_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: Instances_Update_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void instancesUpdateMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        Instance resource = manager.instances().getWithResponse("rgconfigurationmanager", "testname", "testname",
            "testname", com.azure.core.util.Context.NONE).getValue();
        resource.update()
            .withProperties(new InstanceProperties()
                .withSolutionVersionId("vrpzlamkvanqibtjarpxit").withTargetId("tqkdvc")
                .withActiveState(ActiveState.ACTIVE).withReconciliationPolicy(new ReconciliationPolicyProperties()
                    .withState(ReconciliationState.INACTIVE).withInterval("cmzlrjwnlshnkgv"))
                .withSolutionScope("testname"))
            .apply();
    }
}
