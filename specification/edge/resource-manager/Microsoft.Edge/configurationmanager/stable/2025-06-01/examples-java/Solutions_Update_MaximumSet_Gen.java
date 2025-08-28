
import com.azure.resourcemanager.workloadorchestration.models.Solution;
import com.azure.resourcemanager.workloadorchestration.models.SolutionProperties;

/**
 * Samples for Solutions Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Solutions_Update_MaximumSet_Gen.json
     */
    /**
     * Sample code: Solutions_Update_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void solutionsUpdateMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        Solution resource = manager.solutions()
            .getWithResponse("rgconfigurationmanager", "testname", "testname", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withProperties(new SolutionProperties()).apply();
    }
}
