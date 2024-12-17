
/**
 * Samples for DataflowEndpoint Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/DataflowEndpoint_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataflowEndpoint_Get.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void dataflowEndpointGet(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowEndpoints().getWithResponse("rgiotoperations", "resource-name123", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
