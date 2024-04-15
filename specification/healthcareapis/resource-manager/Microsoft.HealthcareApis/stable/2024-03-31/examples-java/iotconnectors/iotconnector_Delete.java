
/**
 * Samples for IotConnectors Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/iotconnectors/
     * iotconnector_Delete.json
     */
    /**
     * Sample code: Delete an IoT Connector.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void deleteAnIoTConnector(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.iotConnectors().delete("testRG", "blue", "workspace1", com.azure.core.util.Context.NONE);
    }
}
