import com.azure.core.util.Context;

/** Samples for EnvironmentVersions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/EnvironmentVersion/get.json
     */
    /**
     * Sample code: Get Environment Version.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getEnvironmentVersion(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.environmentVersions().getWithResponse("test-rg", "my-aml-workspace", "string", "string", Context.NONE);
    }
}
