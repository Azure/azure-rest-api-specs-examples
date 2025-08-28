
import com.azure.resourcemanager.workloadorchestration.models.SolutionDependencyParameter;
import com.azure.resourcemanager.workloadorchestration.models.SolutionTemplateParameter;
import java.util.Arrays;

/**
 * Samples for Targets ResolveConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/Targets_ResolveConfiguration_MaximumSet_Gen.json
     */
    /**
     * Sample code: Targets_ResolveConfiguration_MaximumSet.
     * 
     * @param manager Entry point to WorkloadOrchestrationManager.
     */
    public static void targetsResolveConfigurationMaximumSet(
        com.azure.resourcemanager.workloadorchestration.WorkloadOrchestrationManager manager) {
        manager.targets().resolveConfiguration("rgconfigurationmanager", "testname", new SolutionTemplateParameter()
            .withSolutionTemplateVersionId(
                "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/{resourceProviderNamespace}/{resourceType}/{resourceName}/{resourceType}/{resourceName}")
            .withSolutionInstanceName("testname")
            .withSolutionDependencies(Arrays.asList(new SolutionDependencyParameter()
                .withSolutionVersionId("cydzqntmjlqtksbavjwteru").withSolutionTemplateId("liqauthxnscodbiwktwfwrrsg")
                .withSolutionTemplateVersion("gordjasyxxrj").withSolutionInstanceName("testname")
                .withTargetId("steadvphxtyhjokqicrtg").withDependencies(Arrays.asList()))),
            com.azure.core.util.Context.NONE);
    }
}
