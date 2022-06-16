import com.azure.core.util.Context;

/** Samples for ComponentVersions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/ComponentVersion/get.json
     */
    /**
     * Sample code: Get Component Version.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getComponentVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.componentVersions().getWithResponse("test-rg", "my-aml-workspace", "string", "string", Context.NONE);
    }
}
