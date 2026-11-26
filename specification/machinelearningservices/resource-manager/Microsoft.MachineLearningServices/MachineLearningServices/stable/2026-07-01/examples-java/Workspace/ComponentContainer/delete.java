
/**
 * Samples for ComponentContainers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/ComponentContainer/delete.json
     */
    /**
     * Sample code: Delete Workspace Component Container.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteWorkspaceComponentContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.componentContainers().deleteWithResponse("test-rg", "my-aml-workspace", "string",
            com.azure.core.util.Context.NONE);
    }
}
