
/**
 * Samples for ComponentContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/ComponentContainer/list.json
     */
    /**
     * Sample code: List Workspace Component Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listWorkspaceComponentContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.componentContainers().list("test-rg", "my-aml-workspace", null, null, com.azure.core.util.Context.NONE);
    }
}
