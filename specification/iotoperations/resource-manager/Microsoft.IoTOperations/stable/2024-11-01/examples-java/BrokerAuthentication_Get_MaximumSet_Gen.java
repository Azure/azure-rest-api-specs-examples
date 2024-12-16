
/**
 * Samples for BrokerAuthentication Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/BrokerAuthentication_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: BrokerAuthentication_Get.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void brokerAuthenticationGet(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokerAuthentications().getWithResponse("rgiotoperations", "resource-name123", "resource-name123",
            "resource-name123", com.azure.core.util.Context.NONE);
    }
}
