
/**
 * Samples for CodeVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/CodeVersion/get.json
     */
    /**
     * Sample code: Get Workspace Code Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspaceCodeVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.codeVersions().getWithResponse("test-rg", "my-aml-workspace", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}
