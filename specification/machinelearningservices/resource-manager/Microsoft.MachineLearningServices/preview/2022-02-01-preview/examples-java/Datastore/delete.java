import com.azure.core.util.Context;

/** Samples for Datastores Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Datastore/delete.json
     */
    /**
     * Sample code: Delete datastore.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void deleteDatastore(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager.datastores().deleteWithResponse("test-rg", "my-aml-workspace", "string", Context.NONE);
    }
}
