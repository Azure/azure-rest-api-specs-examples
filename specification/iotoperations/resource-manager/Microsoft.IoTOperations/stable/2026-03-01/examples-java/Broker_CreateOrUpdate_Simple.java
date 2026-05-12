
import com.azure.resourcemanager.iotoperations.models.BackendChain;
import com.azure.resourcemanager.iotoperations.models.BrokerMemoryProfile;
import com.azure.resourcemanager.iotoperations.models.BrokerProperties;
import com.azure.resourcemanager.iotoperations.models.Cardinality;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;
import com.azure.resourcemanager.iotoperations.models.Frontend;
import com.azure.resourcemanager.iotoperations.models.GenerateResourceLimits;
import com.azure.resourcemanager.iotoperations.models.OperationalMode;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Broker CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/Broker_CreateOrUpdate_Simple.json
     */
    /**
     * Sample code: Broker_CreateOrUpdate_Simple.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        brokerCreateOrUpdateSimple(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager
            .brokers().define("resource-name123").withExistingInstance("rgiotoperations",
                "resource-name123")
            .withProperties(new BrokerProperties()
                .withCardinality(new Cardinality()
                    .withBackendChain(new BackendChain().withPartitions(2).withRedundancyFactor(2).withWorkers(2))
                    .withFrontend(new Frontend().withReplicas(2).withWorkers(2)))
                .withGenerateResourceLimits(new GenerateResourceLimits().withCpu(OperationalMode.ENABLED))
                .withMemoryProfile(BrokerMemoryProfile.LOW))
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
