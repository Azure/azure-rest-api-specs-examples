
/**
 * Samples for Datastores Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Datastore/delete.json
     */
    /**
     * Sample code: Delete datastore.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void deleteDatastore(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.datastores().deleteWithResponse("test-rg", "my-aml-workspace", "string",
            com.azure.core.util.Context.NONE);
    }
}
