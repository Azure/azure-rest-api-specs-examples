
/**
 * Samples for IotConnectors ListByWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/iotconnectors/
     * iotconnector_List.json
     */
    /**
     * Sample code: List iotconnectors.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void listIotconnectors(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.iotConnectors().listByWorkspace("testRG", "workspace1", com.azure.core.util.Context.NONE);
    }
}
