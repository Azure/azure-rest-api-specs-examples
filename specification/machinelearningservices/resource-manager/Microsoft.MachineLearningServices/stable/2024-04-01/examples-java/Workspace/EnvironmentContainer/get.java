
/**
 * Samples for EnvironmentContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/EnvironmentContainer/get.json
     */
    /**
     * Sample code: Get Workspace Environment Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspaceEnvironmentContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.environmentContainers().getWithResponse("testrg123", "testworkspace", "testEnvironment",
            com.azure.core.util.Context.NONE);
    }
}
