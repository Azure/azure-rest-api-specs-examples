import com.azure.core.util.Context;

/** Samples for EnvironmentVersions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/EnvironmentVersion/list.json
     */
    /**
     * Sample code: List Environment Version.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listEnvironmentVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .environmentVersions()
            .list("test-rg", "my-aml-workspace", "string", "string", 1, null, null, Context.NONE);
    }
}
