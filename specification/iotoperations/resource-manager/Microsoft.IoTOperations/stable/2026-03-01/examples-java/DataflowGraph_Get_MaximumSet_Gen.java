
/**
 * Samples for DataflowGraph Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/DataflowGraph_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataflowGraph_Get_MaximumSet.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowGraphGetMaximumSet(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowGraphs().getWithResponse("rgiotoperations", "resource-123", "resource-123", "resource-123",
            com.azure.core.util.Context.NONE);
    }
}
