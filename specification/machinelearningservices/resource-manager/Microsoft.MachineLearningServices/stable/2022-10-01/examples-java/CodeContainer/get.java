/** Samples for CodeContainers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/CodeContainer/get.json
     */
    /**
     * Sample code: Get Code Container.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getCodeContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager
            .codeContainers()
            .getWithResponse("testrg123", "testworkspace", "testContainer", com.azure.core.util.Context.NONE);
    }
}
