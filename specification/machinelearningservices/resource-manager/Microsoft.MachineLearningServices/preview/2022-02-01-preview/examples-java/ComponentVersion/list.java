import com.azure.core.util.Context;

/** Samples for ComponentVersions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/ComponentVersion/list.json
     */
    /**
     * Sample code: List Component Version.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listComponentVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .componentVersions()
            .list("test-rg", "my-aml-workspace", "string", "string", 1, null, null, Context.NONE);
    }
}
