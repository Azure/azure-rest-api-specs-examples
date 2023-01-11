import com.azure.core.util.Context;

/** Samples for CodeContainers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/CodeContainer/list.json
     */
    /**
     * Sample code: List Code Container.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listCodeContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.codeContainers().list("testrg123", "testworkspace", null, Context.NONE);
    }
}
