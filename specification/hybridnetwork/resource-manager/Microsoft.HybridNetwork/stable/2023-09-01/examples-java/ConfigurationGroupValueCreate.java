
import com.azure.resourcemanager.hybridnetwork.models.ConfigurationValueWithoutSecrets;
import com.azure.resourcemanager.hybridnetwork.models.OpenDeploymentResourceReference;

/**
 * Samples for ConfigurationGroupValues CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hybridnetwork/resource-manager/Microsoft.HybridNetwork/stable/2023-09-01/examples/
     * ConfigurationGroupValueCreate.json
     */
    /**
     * Sample code: Create or update configuration group value.
     * 
     * @param manager Entry point to HybridNetworkManager.
     */
    public static void
        createOrUpdateConfigurationGroupValue(com.azure.resourcemanager.hybridnetwork.HybridNetworkManager manager) {
        manager.configurationGroupValues().define("testConfigurationGroupValue").withRegion("eastus")
            .withExistingResourceGroup("rg1")
            .withProperties(new ConfigurationValueWithoutSecrets()
                .withConfigurationGroupSchemaResourceReference(new OpenDeploymentResourceReference().withId(
                    "/subscriptions/subid/resourcegroups/testRG/providers/microsoft.hybridnetwork/publishers/testPublisher/configurationGroupSchemas/testConfigurationGroupSchemaName"))
                .withConfigurationValue(
                    "{\"interconnect-groups\":{\"stripe-one\":{\"name\":\"Stripe one\",\"international-interconnects\":[\"france\",\"germany\"],\"domestic-interconnects\":[\"birmingham\",\"edinburgh\"]},\"stripe-two\":{\"name\":\"Stripe two\",\"international-interconnects\":[\"germany\",\"italy\"],\"domestic-interconnects\":[\"edinburgh\",\"london\"]}},\"interconnect-group-assignments\":{\"ssc-one\":{\"ssc\":\"SSC 1\",\"interconnects\":\"stripe-one\"},\"ssc-two\":{\"ssc\":\"SSC 2\",\"interconnects\":\"stripe-two\"}}}"))
            .create();
    }
}
