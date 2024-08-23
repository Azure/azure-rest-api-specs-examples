
/**
 * Samples for Datastores ListSecrets.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Datastore/listSecrets.json
     */
    /**
     * Sample code: Get datastore secrets.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void getDatastoreSecrets(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.datastores().listSecretsWithResponse("test-rg", "my-aml-workspace", "string",
            com.azure.core.util.Context.NONE);
    }
}
