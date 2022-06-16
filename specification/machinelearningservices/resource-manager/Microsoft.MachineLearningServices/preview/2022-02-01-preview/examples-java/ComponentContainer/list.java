import com.azure.core.util.Context;

/** Samples for ComponentContainers List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/ComponentContainer/list.json
     */
    /**
     * Sample code: List Component Container.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listComponentContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.componentContainers().list("test-rg", "my-aml-workspace", null, null, Context.NONE);
    }
}
