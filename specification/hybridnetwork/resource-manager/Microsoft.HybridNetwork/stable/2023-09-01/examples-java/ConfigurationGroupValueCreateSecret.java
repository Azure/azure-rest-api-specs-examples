
import com.azure.resourcemanager.hybridnetwork.models.ConfigurationValueWithSecrets;
import com.azure.resourcemanager.hybridnetwork.models.OpenDeploymentResourceReference;

/**
 * Samples for ConfigurationGroupValues CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ConfigurationGroupValueCreateSecret.json
     */
    /**
     * Sample code: Create or update configuration group value with secrets.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void createOrUpdateConfigurationGroupValueWithSecrets(
        com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.configurationGroupValues().define("testConfigurationGroupValue").withRegion("eastus")
            .withExistingResourceGroup("rg1")
            .withProperties(new ConfigurationValueWithSecrets()
                .withConfigurationGroupSchemaResourceReference(new OpenDeploymentResourceReference().withId(
                    "/subscriptions/subid/resourcegroups/testRG/providers/microsoft.hybridnetwork/publishers/testPublisher/configurationGroupSchemas/testConfigurationGroupSchemaName"))
                .withSecretConfigurationValue("fakeTokenPlaceholder"))
            .create();
    }
}
