/** Samples for CodeContainers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/CodeContainer/delete.json
     */
    /**
     * Sample code: Delete Code Container.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteCodeContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .codeContainers()
            .deleteWithResponse("testrg123", "testworkspace", "testContainer", com.azure.core.util.Context.NONE);
    }
}
