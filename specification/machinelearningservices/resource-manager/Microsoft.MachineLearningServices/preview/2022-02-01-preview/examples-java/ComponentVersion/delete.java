import com.azure.core.util.Context;

/** Samples for ComponentVersions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/ComponentVersion/delete.json
     */
    /**
     * Sample code: Delete Component Version.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteComponentVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.componentVersions().deleteWithResponse("test-rg", "my-aml-workspace", "string", "string", Context.NONE);
    }
}
