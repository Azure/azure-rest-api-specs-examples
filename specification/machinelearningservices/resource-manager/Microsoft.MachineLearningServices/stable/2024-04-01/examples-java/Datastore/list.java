
import java.util.Arrays;

/**
 * Samples for Datastores List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Datastore/list.json
     */
    /**
     * Sample code: List datastores.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void listDatastores(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.datastores().list("test-rg", "my-aml-workspace", null, 1, false, Arrays.asList("string"), "string",
            "string", false, com.azure.core.util.Context.NONE);
    }
}
