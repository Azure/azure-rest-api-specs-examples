import com.azure.core.util.Context;

/** Samples for FhirDestinations ListByIotConnector. */
public final class Main {
    /*
     * x-ms-original-file: specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2021-11-01/examples/iotconnectors/iotconnector_fhirdestination_List.json
     */
    /**
     * Sample code: List IoT Connectors.
     *
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void listIoTConnectors(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.fhirDestinations().listByIotConnector("testRG", "workspace1", "blue", Context.NONE);
    }
}
