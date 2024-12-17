
import com.azure.resourcemanager.iotoperations.models.AuthorizationConfig;
import com.azure.resourcemanager.iotoperations.models.AuthorizationRule;
import com.azure.resourcemanager.iotoperations.models.BrokerAuthorizationProperties;
import com.azure.resourcemanager.iotoperations.models.BrokerResourceDefinitionMethods;
import com.azure.resourcemanager.iotoperations.models.BrokerResourceRule;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;
import com.azure.resourcemanager.iotoperations.models.OperationalMode;
import com.azure.resourcemanager.iotoperations.models.PrincipalDefinition;
import com.azure.resourcemanager.iotoperations.models.StateStoreResourceDefinitionMethods;
import com.azure.resourcemanager.iotoperations.models.StateStoreResourceKeyTypes;
import com.azure.resourcemanager.iotoperations.models.StateStoreResourceRule;
import java.util.Arrays;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for BrokerAuthorization CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/BrokerAuthorization_CreateOrUpdate_Simple.json
     */
    /**
     * Sample code: BrokerAuthorization_CreateOrUpdate_Simple.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        brokerAuthorizationCreateOrUpdateSimple(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokerAuthorizations().define("resource-name123")
            .withExistingBroker("rgiotoperations", "resource-name123", "resource-name123")
            .withExtendedLocation(
                new ExtendedLocation().withName("qmbrfwcpwwhggszhrdjv").withType(ExtendedLocationType.CUSTOM_LOCATION))
            .withProperties(new BrokerAuthorizationProperties()
                .withAuthorizationPolicies(new AuthorizationConfig().withCache(OperationalMode.ENABLED)
                    .withRules(Arrays.asList(new AuthorizationRule()
                        .withBrokerResources(
                            Arrays.asList(new BrokerResourceRule().withMethod(BrokerResourceDefinitionMethods.CONNECT),
                                new BrokerResourceRule().withMethod(BrokerResourceDefinitionMethods.SUBSCRIBE)
                                    .withTopics(Arrays.asList("topic", "topic/with/wildcard/#"))))
                        .withPrincipals(new PrincipalDefinition()
                            .withAttributes(Arrays.asList(mapOf("floor", "floor1", "site", "site1")))
                            .withClientIds(Arrays.asList("my-client-id")))
                        .withStateStoreResources(Arrays.asList(new StateStoreResourceRule()
                            .withKeyType(StateStoreResourceKeyTypes.PATTERN).withKeys(Arrays.asList("*"))
                            .withMethod(StateStoreResourceDefinitionMethods.READ_WRITE)))))))
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
