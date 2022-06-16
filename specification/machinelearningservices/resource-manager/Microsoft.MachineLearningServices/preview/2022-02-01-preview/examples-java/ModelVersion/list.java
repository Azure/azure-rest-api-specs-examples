import com.azure.core.util.Context;

/** Samples for ModelVersions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/ModelVersion/list.json
     */
    /**
     * Sample code: List Model Version.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listModelVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .modelVersions()
            .list(
                "test-rg",
                "my-aml-workspace",
                "string",
                null,
                "string",
                1,
                "string",
                "string",
                1,
                "string",
                "string",
                null,
                null,
                Context.NONE);
    }
}
