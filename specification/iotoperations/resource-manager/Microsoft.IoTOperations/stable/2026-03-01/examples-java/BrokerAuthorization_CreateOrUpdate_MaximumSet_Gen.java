
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
     * x-ms-original-file: 2026-03-01/BrokerAuthorization_CreateOrUpdate_MaximumSet_Gen.json
     */
    /**
     * Sample code: BrokerAuthorization_CreateOrUpdate.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        brokerAuthorizationCreateOrUpdate(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokerAuthorizations().define("resource-name123")
            .withExistingBroker("rgiotoperations", "resource-name123", "resource-name123")
            .withProperties(new BrokerAuthorizationProperties().withAuthorizationPolicies(new AuthorizationConfig()
                .withCache(OperationalMode.ENABLED)
                .withRules(Arrays.asList(new AuthorizationRule()
                    .withBrokerResources(
                        Arrays.asList(new BrokerResourceRule().withMethod(BrokerResourceDefinitionMethods.CONNECT)
                            .withClientIds(Arrays.asList("nlc")).withTopics(Arrays.asList("wvuca"))))
                    .withPrincipals(new PrincipalDefinition()
                        .withAttributes(Arrays.asList(mapOf("key5526", "fakeTokenPlaceholder")))
                        .withClientIds(Arrays.asList("smopeaeddsygz"))
                        .withUsernames(Arrays.asList("iozngyqndrteikszkbasinzdjtm")))
                    .withStateStoreResources(
                        Arrays.asList(new StateStoreResourceRule().withKeyType(StateStoreResourceKeyTypes.PATTERN)
                            .withKeys(Arrays.asList("tkounsqtwvzyaklxjqoerpu"))
                            .withMethod(StateStoreResourceDefinitionMethods.READ)))))))
            .withExtendedLocation(new ExtendedLocation().withName(
                "/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123")
                .withType(ExtendedLocationType.CUSTOM_LOCATION))
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
