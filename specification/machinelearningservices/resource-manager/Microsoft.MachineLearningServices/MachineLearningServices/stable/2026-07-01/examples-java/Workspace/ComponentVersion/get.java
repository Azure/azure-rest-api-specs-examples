
/**
 * Samples for ComponentVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/ComponentVersion/get.json
     */
    /**
     * Sample code: Get Workspace Component Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspaceComponentVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.componentVersions().getWithResponse("test-rg", "my-aml-workspace", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}
