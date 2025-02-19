
import com.azure.core.management.serializer.SerializerFactory;
import com.azure.core.util.serializer.SerializerEncoding;
import com.azure.resourcemanager.customerinsights.models.ConnectorTypes;
import java.io.IOException;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Connectors CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/customer-insights/resource-manager/Microsoft.CustomerInsights/stable/2017-04-26/examples/
     * ConnectorsCreateOrUpdate.json
     */
    /**
     * Sample code: Connectors_CreateOrUpdate.
     * 
     * @param manager Entry point to CustomerInsightsManager.
     */
    public static void connectorsCreateOrUpdate(
        com.azure.resourcemanager.customerinsights.CustomerInsightsManager manager) throws IOException {
        manager.connectors().define("testConnector").withExistingHub("TestHubRG", "sdkTestHub")
            .withConnectorType(ConnectorTypes.AZURE_BLOB).withDisplayName("testConnector")
            .withDescription("Test connector")
            .withConnectorProperties(mapOf("connectionKeyVaultUrl",
                SerializerFactory.createDefaultManagementSerializerAdapter().deserialize(
                    "{\"organizationId\":\"XXX\",\"organizationUrl\":\"https://XXX.crmlivetie.com/\"}", Object.class,
                    SerializerEncoding.JSON)))
            .create();
    }

    // Use "Map.of" if available
    @SuppressWarnings("unchecked")
    private static <T> Map<String, T> mapOf(Object... inputs) {
        Map<String, T> map = new HashMap<>();
        for (int i = 0; i < inputs.length; i += 2) {
            String key = (String) inputs[i];
            T value = (T) inputs[i + 1];
            map.put(key, value);
        }
        return map;
    }
}
