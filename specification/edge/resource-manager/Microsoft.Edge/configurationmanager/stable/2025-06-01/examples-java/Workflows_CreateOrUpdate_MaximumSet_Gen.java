
import com.azure.resourcemanager.workloadorchestration.models.ExtendedLocation;
import com.azure.resourcemanager.workloadorchestration.models.ExtendedLocationType;
import com.azure.resourcemanager.workloadorchestration.models.WorkflowProperties;

/**
 * Samples for Workflows CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Workflows_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: Workflows_CreateOrUpdate_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void workflowsCreateOrUpdateMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.workflows().define("testname").withExistingContext("rgconfigurationmanager", "testname")
            .withProperties(new WorkflowProperties()).withExtendedLocation(
                new ExtendedLocation().withName("szjrwimeqyiue").withType(ExtendedLocationType.EDGE_ZONE))
            .create();
    }
}
