
/**
 * Samples for ModelDeployments Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-02-preview/ModelDeployments_Get.json
     */
    /**
     * Sample code: ModelDeployments_Get_MaximumSet.
     * 
     * @param manager Entry point to ContainerServiceAIManagerManager.
     */
    public static void modelDeploymentsGetMaximumSet(
        com.azure.resourcemanager.containerserviceaimanager.ContainerServiceAIManagerManager manager) {
        manager.modelDeployments().getWithResponse("rgaimanagers", "aimanager1", "namespace-1", "deployment-1",
            com.azure.core.util.Context.NONE);
    }
}
