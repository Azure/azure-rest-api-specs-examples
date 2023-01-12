/** Samples for ModelVersions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/ModelVersion/get.json
     */
    /**
     * Sample code: Get Model Version.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getModelVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .modelVersions()
            .getWithResponse("test-rg", "my-aml-workspace", "string", "string", com.azure.core.util.Context.NONE);
    }
}
