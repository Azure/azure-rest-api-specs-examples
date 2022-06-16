import com.azure.core.util.Context;

/** Samples for ModelVersions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/ModelVersion/get.json
     */
    /**
     * Sample code: Get Model Version.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getModelVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.modelVersions().getWithResponse("test-rg", "my-aml-workspace", "string", "string", Context.NONE);
    }
}
