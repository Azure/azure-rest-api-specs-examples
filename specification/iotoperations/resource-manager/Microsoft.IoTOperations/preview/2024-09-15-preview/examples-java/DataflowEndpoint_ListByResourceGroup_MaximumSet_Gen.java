
/**
 * Samples for DataflowEndpoint ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-15-preview/DataflowEndpoint_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataflowEndpoint_ListByResourceGroup.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowEndpointListByResourceGroup(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowEndpoints().listByResourceGroup("rgiotoperations", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
