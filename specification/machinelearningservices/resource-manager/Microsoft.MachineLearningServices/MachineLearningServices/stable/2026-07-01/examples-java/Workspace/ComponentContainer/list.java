
/**
 * Samples for ComponentContainers List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/ComponentContainer/list.json
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
