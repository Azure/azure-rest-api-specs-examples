
/**
 * Samples for DataVersions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/DataVersionBase/delete.json
     */
    /**
     * Sample code: Delete Workspace Data Version Base.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        deleteWorkspaceDataVersionBase(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.dataVersions().deleteWithResponse("test-rg", "my-aml-workspace", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}
