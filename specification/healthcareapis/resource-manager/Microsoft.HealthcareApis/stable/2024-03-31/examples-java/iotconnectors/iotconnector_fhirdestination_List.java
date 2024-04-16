
/**
 * Samples for FhirDestinations ListByIotConnector.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/iotconnectors/
     * iotconnector_fhirdestination_List.json
     */
    /**
     * Sample code: List IoT Connectors.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void listIoTConnectors(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.fhirDestinations().listByIotConnector("testRG", "workspace1", "blue", com.azure.core.util.Context.NONE);
    }
}
