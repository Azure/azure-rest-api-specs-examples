
/**
 * Samples for ComponentContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/ComponentContainer/get.json
     */
    /**
     * Sample code: Get Workspace Component Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspaceComponentContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.componentContainers().getWithResponse("test-rg", "my-aml-workspace", "string",
            com.azure.core.util.Context.NONE);
    }
}
