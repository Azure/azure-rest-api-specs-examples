
import com.azure.resourcemanager.datafactory.models.EnableInteractiveQueryRequest;

/**
 * Samples for IntegrationRuntimeOperation EnableInteractiveQuery.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/IntegrationRuntimes_EnableInteractiveQuery.json
     */
    /**
     * Sample code: IntegrationRuntime_EnableInteractiveQuery.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void
        integrationRuntimeEnableInteractiveQuery(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimeOperations().enableInteractiveQuery("exampleResourceGroup", "exampleFactoryName",
            "exampleIntegrationRuntime", new EnableInteractiveQueryRequest().withAutoTerminationMinutes(10),
            com.azure.core.util.Context.NONE);
    }
}
