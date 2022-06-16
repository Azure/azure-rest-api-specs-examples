import com.azure.core.util.Context;

/** Samples for CodeContainers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/CodeContainer/get.json
     */
    /**
     * Sample code: Get Code Container.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getCodeContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.codeContainers().getWithResponse("testrg123", "testworkspace", "testContainer", Context.NONE);
    }
}
