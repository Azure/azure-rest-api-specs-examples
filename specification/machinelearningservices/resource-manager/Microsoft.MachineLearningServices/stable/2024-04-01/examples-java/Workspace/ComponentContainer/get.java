
/**
 * Samples for ComponentContainers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/ComponentContainer/get.json
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
