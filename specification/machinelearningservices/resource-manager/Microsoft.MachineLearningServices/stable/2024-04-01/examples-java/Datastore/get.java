
/**
 * Samples for Datastores Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Datastore/get.json
     */
    /**
     * Sample code: Get datastore.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getDatastore(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.datastores().getWithResponse("test-rg", "my-aml-workspace", "string", com.azure.core.util.Context.NONE);
    }
}
