import com.azure.core.util.Context;

/** Samples for IotConnectorFhirDestination Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/iotconnectors/iotconnector_fhirdestination_Delete.json
     */
    /**
     * Sample code: Delete an IoT Connector destination.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void deleteAnIoTConnectorDestination(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.iotConnectorFhirDestinations().delete("testRG", "workspace1", "blue", "dest1", Context.NONE);
    }
}
