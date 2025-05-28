
import com.azure.resourcemanager.datafactory.models.IntegrationRuntimeAutoUpdate;
import com.azure.resourcemanager.datafactory.models.IntegrationRuntimeResource;

/**
 * Samples for IntegrationRuntimes Update.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * IntegrationRuntimes_Update.json
     */
    /**
     * Sample code: IntegrationRuntimes_Update.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void integrationRuntimesUpdate(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        IntegrationRuntimeResource resource = manager.integrationRuntimes().getWithResponse("exampleResourceGroup",
            "exampleFactoryName", "exampleIntegrationRuntime", null, com.azure.core.util.Context.NONE).getValue();
        resource.update().withAutoUpdate(IntegrationRuntimeAutoUpdate.OFF).withUpdateDelayOffset("\"PT3H\"").apply();
    }
}
