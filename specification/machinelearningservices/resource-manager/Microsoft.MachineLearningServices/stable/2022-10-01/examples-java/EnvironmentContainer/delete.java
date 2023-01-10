import com.azure.core.util.Context;

/** Samples for EnvironmentContainers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/EnvironmentContainer/delete.json
     */
    /**
     * Sample code: Delete Environment Container.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteEnvironmentContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.environmentContainers().deleteWithResponse("testrg123", "testworkspace", "testContainer", Context.NONE);
    }
}
