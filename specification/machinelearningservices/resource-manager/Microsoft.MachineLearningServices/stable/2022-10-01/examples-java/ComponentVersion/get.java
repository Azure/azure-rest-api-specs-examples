/** Samples for ComponentVersions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/ComponentVersion/get.json
     */
    /**
     * Sample code: Get Component Version.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getComponentVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .componentVersions()
            .getWithResponse("test-rg", "my-aml-workspace", "string", "string", com.azure.core.util.Context.NONE);
    }
}
