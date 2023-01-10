import com.azure.core.util.Context;

/** Samples for EnvironmentContainers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/EnvironmentContainer/list.json
     */
    /**
     * Sample code: List Environment Container.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listEnvironmentContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.environmentContainers().list("testrg123", "testworkspace", null, null, Context.NONE);
    }
}
