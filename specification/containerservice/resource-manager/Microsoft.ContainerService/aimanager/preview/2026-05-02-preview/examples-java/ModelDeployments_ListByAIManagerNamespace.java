
/**
 * Samples for ModelDeployments ListByAIManagerNamespace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/ModelDeployments_ListByAIManagerNamespace.json
     */
    /**
     * Sample code: ModelDeployments_ListByAIManagerNamespace_MaximumSet.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void modelDeploymentsListByAIManagerNamespaceMaximumSet(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.modelDeployments().listByAIManagerNamespace("rgaimanagers", "aimanager1", "namespace-1",
            com.azure.core.util.Context.NONE);
    }
}
