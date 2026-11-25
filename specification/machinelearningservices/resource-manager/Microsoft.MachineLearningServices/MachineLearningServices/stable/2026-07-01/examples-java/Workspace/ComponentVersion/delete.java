
/**
 * Samples for ComponentVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/ComponentVersion/delete.json
     */
    /**
     * Sample code: Delete Workspace Component Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteWorkspaceComponentVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.componentVersions().deleteWithResponse("test-rg", "my-aml-workspace", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}
