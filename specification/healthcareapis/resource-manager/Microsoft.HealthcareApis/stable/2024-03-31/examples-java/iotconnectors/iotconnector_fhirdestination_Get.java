
/**
 * Samples for IotConnectorFhirDestination Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/iotconnectors/
     * iotconnector_fhirdestination_Get.json
     */
    /**
     * Sample code: Get an IoT Connector destination.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void
        getAnIoTConnectorDestination(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.iotConnectorFhirDestinations().getWithResponse("testRG", "workspace1", "blue", "dest1",
            com.azure.core.util.Context.NONE);
    }
}
