/** Samples for ModelVersions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/ModelVersion/delete.json
     */
    /**
     * Sample code: Delete Model Version.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteModelVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .modelVersions()
            .deleteWithResponse("test-rg", "my-aml-workspace", "string", "string", com.azure.core.util.Context.NONE);
    }
}
