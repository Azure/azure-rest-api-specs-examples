
import com.azure.resourcemanager.workloadorchestration.models.InstallSolutionParameter;

/**
 * Samples for Targets InstallSolution.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Targets_InstallSolution_MaximumSet_Gen.json
     */
    /**
     * Sample code: Targets_InstallSolution_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void targetsInstallSolutionMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.targets().installSolution("rgconfigurationmanager", "testname",
            new InstallSolutionParameter().withSolutionVersionId(
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
            com.azure.core.util.Context.NONE);
    }
}
