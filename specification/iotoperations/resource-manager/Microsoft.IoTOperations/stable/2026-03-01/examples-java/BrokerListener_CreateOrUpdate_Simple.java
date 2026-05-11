
import com.azure.resourcemanager.iotoperations.models.BrokerListenerProperties;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocation;
import com.azure.resourcemanager.iotoperations.models.ExtendedLocationType;
import com.azure.resourcemanager.iotoperations.models.ListenerPort;
import java.util.Arrays;

/**
 * Samples for BrokerListener CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/BrokerListener_CreateOrUpdate_Simple.json
     */
    /**
     * Sample code: BrokerListener_CreateOrUpdate_Simple.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        brokerListenerCreateOrUpdateSimple(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokerListeners().define("resource-name123")
            .withExistingBroker("rgiotoperations", "resource-name123", "resource-name123")
            .withProperties(new BrokerListenerProperties().withPorts(Arrays.asList(new ListenerPort().withPort(1883))))
            .withExtendedLocation(new ExtendedLocation().withName(
                "/subscriptions/F8C729F9-DF9C-4743-848F-96EE433D8E53/resourceGroups/rgiotoperations/providers/Microsoft.ExtendedLocation/customLocations/resource-123")
                .withType(ExtendedLocationType.CUSTOM_LOCATION))
            .create();
    }
}
