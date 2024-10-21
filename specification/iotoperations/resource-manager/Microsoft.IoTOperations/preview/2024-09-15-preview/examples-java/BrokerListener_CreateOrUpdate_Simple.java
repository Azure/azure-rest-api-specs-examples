
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
     * x-ms-original-file: 2024-09-15-preview/BrokerListener_CreateOrUpdate_Simple.json
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
            .withExtendedLocation(
                new ExtendedLocation().withName("qmbrfwcpwwhggszhrdjv").withType(ExtendedLocationType.CUSTOM_LOCATION))
            .withProperties(new BrokerListenerProperties().withPorts(Arrays.asList(new ListenerPort().withPort(1883))))
            .create();
    }
}
