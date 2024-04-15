
/**
 * Samples for IotConnectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/iotconnectors/
     * iotconnector_Get.json
     */
    /**
     * Sample code: Get an IoT Connector.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void getAnIoTConnector(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.iotConnectors().getWithResponse("testRG", "workspace1", "blue", com.azure.core.util.Context.NONE);
    }
}
