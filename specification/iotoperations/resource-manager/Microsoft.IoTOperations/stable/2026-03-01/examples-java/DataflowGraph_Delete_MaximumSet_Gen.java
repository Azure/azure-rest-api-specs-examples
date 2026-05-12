
/**
 * Samples for DataflowGraph Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/DataflowGraph_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataflowGraph_Delete_MaximumSet.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        dataflowGraphDeleteMaximumSet(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowGraphs().delete("rgiotoperations", "resource-123", "resource-123", "resource-123",
            com.azure.core.util.Context.NONE);
    }
}
