import com.azure.core.util.Context;

/** Samples for Datastores Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Datastore/get.json
     */
    /**
     * Sample code: Get datastore.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void getDatastore(com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.datastores().getWithResponse("test-rg", "my-aml-workspace", "string", Context.NONE);
    }
}
