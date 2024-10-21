
/**
 * Samples for BrokerAuthorization Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-15-preview/BrokerAuthorization_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: BrokerAuthorization_Get.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void brokerAuthorizationGet(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.brokerAuthorizations().getWithResponse("rgiotoperations", "resource-name123", "resource-name123",
            "resource-name123", com.azure.core.util.Context.NONE);
    }
}
