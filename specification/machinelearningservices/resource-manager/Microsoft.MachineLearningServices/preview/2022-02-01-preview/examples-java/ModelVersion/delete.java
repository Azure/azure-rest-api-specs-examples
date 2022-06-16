import com.azure.core.util.Context;

/** Samples for ModelVersions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/ModelVersion/delete.json
     */
    /**
     * Sample code: Delete Model Version.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteModelVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.modelVersions().deleteWithResponse("test-rg", "my-aml-workspace", "string", "string", Context.NONE);
    }
}
