
/**
 * Samples for CodeContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/CodeContainer/get.json
     */
    /**
     * Sample code: Get Workspace Code Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspaceCodeContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.codeContainers().getWithResponse("testrg123", "testworkspace", "testContainer",
            com.azure.core.util.Context.NONE);
    }
}
