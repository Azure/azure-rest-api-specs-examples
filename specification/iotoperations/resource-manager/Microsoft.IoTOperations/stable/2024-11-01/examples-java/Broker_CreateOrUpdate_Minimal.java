
import com.azure.resourcemanager.iotoperations.models.BrokerMemoryProfile;
import com.azure.resourcemanager.iotoperations.models.BrokerProperties;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;
import java.util.HashMap;
import java.util.Map;

/**
 * Samples for Broker CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/Broker_CreateOrUpdate_Minimal.json
     */
    /**
     * Sample code: Broker_CreateOrUpdate_Minimal.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        brokerCreateOrUpdateMinimal(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokers().define("resource-name123").withExistingInstance("rgiotoperations", "resource-name123")
            .withExtendedLocation(
                new ExtendedLocation().withName("qmbrfwcpwwhggszhrdjv").withType(ExtendedLocationType.CUSTOM_LOCATION))
            .withProperties(new BrokerProperties().withMemoryProfile(BrokerMemoryProfile.TINY)).create();
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
