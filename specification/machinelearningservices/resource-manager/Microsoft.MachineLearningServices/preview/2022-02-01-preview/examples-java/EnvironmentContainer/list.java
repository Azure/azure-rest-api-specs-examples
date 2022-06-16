import com.azure.core.util.Context;

/** Samples for EnvironmentContainers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/EnvironmentContainer/list.json
     */
    /**
     * Sample code: List Environment Container.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listEnvironmentContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.environmentContainers().list("testrg123", "testworkspace", null, null, Context.NONE);
    }
}
