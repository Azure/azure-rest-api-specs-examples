import com.azure.core.util.Context;

/** Samples for ComponentContainers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2022-10-01/examples/ComponentContainer/get.json
     */
    /**
     * Sample code: Get Component Container.
     *
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getComponentContainer(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.componentContainers().getWithResponse("test-rg", "my-aml-workspace", "string", Context.NONE);
    }
}
