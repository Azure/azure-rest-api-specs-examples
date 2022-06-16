import com.azure.core.util.Context;

/** Samples for ComponentContainers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/ComponentContainer/get.json
     */
    /**
     * Sample code: Get Component Container.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getComponentContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.componentContainers().getWithResponse("test-rg", "my-aml-workspace", "string", Context.NONE);
    }
}
