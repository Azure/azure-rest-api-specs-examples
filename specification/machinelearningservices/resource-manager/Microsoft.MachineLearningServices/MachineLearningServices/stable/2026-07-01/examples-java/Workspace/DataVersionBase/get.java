
/**
 * Samples for DataVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Workspace/DataVersionBase/get.json
     */
    /**
     * Sample code: Get Workspace Data Version Base.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getWorkspaceDataVersionBase(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.dataVersions().getWithResponse("test-rg", "my-aml-workspace", "string", "string",
            com.azure.core.util.Context.NONE);
    }
}
