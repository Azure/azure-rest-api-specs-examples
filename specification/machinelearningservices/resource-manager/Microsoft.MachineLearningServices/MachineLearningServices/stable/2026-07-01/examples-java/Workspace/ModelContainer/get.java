
/**
 * Samples for ModelContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/ModelContainer/get.json
     */
    /**
     * Sample code: Get Workspace Model Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspaceModelContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.modelContainers().getWithResponse("testrg123", "workspace123", "testContainer",
            com.azure.core.util.Context.NONE);
    }
}
