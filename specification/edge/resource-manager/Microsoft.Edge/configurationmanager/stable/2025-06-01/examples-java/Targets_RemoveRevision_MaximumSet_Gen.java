
import com.azure.resourcemanager.workloadorchestration.models.RemoveRevisionParameter;

/**
 * Samples for Targets RemoveRevision.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Targets_RemoveRevision_MaximumSet_Gen.json
     */
    /**
     * Sample code: Targets_RemoveRevision_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void targetsRemoveRevisionMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.targets().removeRevision("rgconfigurationmanager", "testname",
            new RemoveRevisionParameter().withSolutionTemplateId(
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}")
                .withSolutionVersion("tomwmqybqomwkfaeukjneva"),
            com.azure.core.util.Context.NONE);
    }
}
