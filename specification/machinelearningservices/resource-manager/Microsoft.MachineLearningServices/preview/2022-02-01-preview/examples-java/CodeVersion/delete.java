import com.azure.core.util.Context;

/** Samples for CodeVersions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/CodeVersion/delete.json
     */
    /**
     * Sample code: Delete Code Version.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteCodeVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.codeVersions().deleteWithResponse("test-rg", "my-aml-workspace", "string", "string", Context.NONE);
    }
}
