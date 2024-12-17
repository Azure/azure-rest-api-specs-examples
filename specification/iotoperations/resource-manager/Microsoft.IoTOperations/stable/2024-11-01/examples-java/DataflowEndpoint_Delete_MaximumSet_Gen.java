
/**
 * Samples for DataflowEndpoint Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-11-01/DataflowEndpoint_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataflowEndpoint_Delete.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void dataflowEndpointDelete(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowEndpoints().delete("rgiotoperations", "resource-name123", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
