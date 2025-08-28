
import com.azure.resourcemanager.workloadorchestration.models.UninstallSolutionParameter;

/**
 * Samples for Targets UninstallSolution.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Targets_UninstallSolution_MaximumSet_Gen.json
     */
    /**
     * Sample code: Targets_UninstallSolution_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void targetsUninstallSolutionMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.targets().uninstallSolution("rgconfigurationmanager", "testname",
            new UninstallSolutionParameter().withSolutionTemplateId(
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}")
                .withSolutionInstanceName("lzihiumrcxbolmkqktvtuqyhg"),
            com.azure.core.util.Context.NONE);
    }
}
