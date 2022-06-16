import com.azure.core.util.Context;

/** Samples for ComponentContainers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/ComponentContainer/delete.json
     */
    /**
     * Sample code: Delete Component Container.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteComponentContainer(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.componentContainers().deleteWithResponse("test-rg", "my-aml-workspace", "string", Context.NONE);
    }
}
