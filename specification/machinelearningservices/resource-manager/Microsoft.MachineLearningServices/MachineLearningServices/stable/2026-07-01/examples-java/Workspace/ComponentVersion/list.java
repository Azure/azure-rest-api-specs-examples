
/**
 * Samples for ComponentVersions List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/ComponentVersion/list.json
     */
    /**
     * Sample code: List Workspace Component Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        listWorkspaceComponentVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.componentVersions().list("test-rg", "my-aml-workspace", "string", "string", 1, null, null,
            com.azure.core.util.Context.NONE);
    }
}
