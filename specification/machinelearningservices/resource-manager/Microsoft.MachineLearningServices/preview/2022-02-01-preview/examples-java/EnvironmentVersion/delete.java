import com.azure.core.util.Context;

/** Samples for EnvironmentVersions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/EnvironmentVersion/delete.json
     */
    /**
     * Sample code: Delete Environment Version.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteEnvironmentVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .environmentVersions()
            .deleteWithResponse("test-rg", "my-aml-workspace", "string", "string", Context.NONE);
    }
}
