
/**
 * Samples for ModelDeployments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/ModelDeployments_Delete.json
     */
    /**
     * Sample code: ModelDeployments_Delete_MaximumSet.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void modelDeploymentsDeleteMaximumSet(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.modelDeployments().delete("rgaimanagers", "aimanager1", "namespace-1", "deployment-1",
            "\"abc123def456\"", com.azure.core.util.Context.NONE);
    }
}
