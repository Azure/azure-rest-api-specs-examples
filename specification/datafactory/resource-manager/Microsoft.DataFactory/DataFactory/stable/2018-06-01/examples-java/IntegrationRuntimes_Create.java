
import com.azure.resourcemanager.datafactory.models.SelfHostedIntegrationRuntime;

/**
 * Samples for IntegrationRuntimes CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2018-06-01/IntegrationRuntimes_Create.json
     */
    /**
     * Sample code: IntegrationRuntimes_Create.
     * 
     * @param manager Entry point to DataFactoryManager.
     */
    public static void integrationRuntimesCreate(com.azure.resourcemanager.datafactory.DataFactoryManager manager) {
        manager.integrationRuntimes().define("exampleIntegrationRuntime")
            .withExistingFactory("exampleResourceGroup", "exampleFactoryName")
            .withProperties(new SelfHostedIntegrationRuntime().withDescription("A selfhosted integration runtime"))
            .create();
    }
}
