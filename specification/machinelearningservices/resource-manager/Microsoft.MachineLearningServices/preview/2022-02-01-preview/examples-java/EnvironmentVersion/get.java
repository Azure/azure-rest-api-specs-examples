import com.azure.core.util.Context;

/** Samples for EnvironmentVersions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/EnvironmentVersion/get.json
     */
    /**
     * Sample code: Get Environment Version.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getEnvironmentVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.environmentVersions().getWithResponse("test-rg", "my-aml-workspace", "string", "string", Context.NONE);
    }
}
