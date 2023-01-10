import com.azure.core.util.Context;

/** Samples for ComponentVersions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/ComponentVersion/list.json
     */
    /**
     * Sample code: List Component Version.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listComponentVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .componentVersions()
            .list("test-rg", "my-aml-workspace", "string", "string", 1, null, null, Context.NONE);
    }
}
