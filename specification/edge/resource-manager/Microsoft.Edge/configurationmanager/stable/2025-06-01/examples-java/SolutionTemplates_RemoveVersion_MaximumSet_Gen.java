
import com.azure.resourcemanager.workloadorchestration.models.VersionParameter;

/**
 * Samples for SolutionTemplates RemoveVersion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/SolutionTemplates_RemoveVersion_MaximumSet_Gen.json
     */
    /**
     * Sample code: SolutionTemplates_RemoveVersion_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void solutionTemplatesRemoveVersionMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.solutionTemplates().removeVersion("rgconfigurationmanager", "testname",
            new VersionParameter().withVersion("ghtvdzgmzncaifrnuumg"), com.azure.core.util.Context.NONE);
    }
}
