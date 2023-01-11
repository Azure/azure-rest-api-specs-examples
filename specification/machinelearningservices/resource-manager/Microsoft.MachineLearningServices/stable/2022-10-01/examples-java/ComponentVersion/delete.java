/** Samples for ComponentVersions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/ComponentVersion/delete.json
     */
    /**
     * Sample code: Delete Component Version.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteComponentVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .componentVersions()
            .deleteWithResponse("test-rg", "my-aml-workspace", "string", "string", com.azure.core.util.Context.NONE);
    }
}
