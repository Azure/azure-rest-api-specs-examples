
import com.azure.resourcemanager.datafactory.models.LinkedIntegrationRuntimeRequest;

/**
 * Samples for IntegrationRuntimes RemoveLinks.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/
     * IntegrationRuntimes_RemoveLinks.json
     */
    /**
     * Sample code: IntegrationRuntimes_Upgrade.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void integrationRuntimesUpgrade(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimes().removeLinksWithResponse("exampleResourceGroup", "exampleFactoryName",
            "exampleIntegrationRuntime",
            new LinkedIntegrationRuntimeRequest().withLinkedFactoryName("exampleFactoryName-linked"),
            com.azure.core.util.Context.NONE);
    }
}
