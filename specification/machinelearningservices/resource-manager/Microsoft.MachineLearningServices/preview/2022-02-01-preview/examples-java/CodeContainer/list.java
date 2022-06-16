import com.azure.core.util.Context;

/** Samples for CodeContainers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/CodeContainer/list.json
     */
    /**
     * Sample code: List Code Container.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listCodeContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.codeContainers().list("testrg123", "testworkspace", null, Context.NONE);
    }
}
