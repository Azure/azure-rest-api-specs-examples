
/**
 * Samples for DataVersions Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/DataVersionBase/get.json
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
