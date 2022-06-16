import com.azure.core.util.Context;
import java.util.Arrays;

/** Samples for Datastores List. */
public final class Main {
    /*
     * x-ms-original-file: specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/preview/2022-02-01-preview/examples/Datastore/list.json
     */
    /**
     * Sample code: List datastores.
     *
     * @param manager Entry point to MachineLearningServicesManager.
     */
    public static void listDatastores(
        com.azure.resourcemanager.machinelearning.MachineLearningServicesManager manager) {
        manager
            .datastores()
            .list(
                "test-rg",
                "my-aml-workspace",
                null,
                1,
                false,
                Arrays.asList("string"),
                "string",
                "string",
                false,
                Context.NONE);
    }
}
