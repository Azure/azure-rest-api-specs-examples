
import com.azure.resourcemanager.workloadorchestration.models.SolutionVersionParameter;

/**
 * Samples for Targets PublishSolutionVersion.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Targets_PublishSolutionVersion_MaximumSet_Gen.json
     */
    /**
     * Sample code: Targets_PublishSolutionVersion_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void targetsPublishSolutionVersionMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.targets().publishSolutionVersion("rgconfigurationmanager", "testname",
            new SolutionVersionParameter().withSolutionVersionId(
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}"),
            com.azure.core.util.Context.NONE);
    }
}
