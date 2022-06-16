import com.azure.core.util.Context;

/** Samples for IotConnectorFhirDestination Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/iotconnectors/iotconnector_fhirdestination_Get.json
     */
    /**
     * Sample code: Get an IoT Connector destination.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void getAnIoTConnectorDestination(
        com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.iotConnectorFhirDestinations().getWithResponse("testRG", "workspace1", "blue", "dest1", Context.NONE);
    }
}
