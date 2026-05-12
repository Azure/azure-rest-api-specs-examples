
/**
 * Samples for DataflowGraph ListByDataflowProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/DataflowGraph_ListByDataflowProfile_MaximumSet_Gen.json
     */
    /**
     * Sample code: DataflowGraph_ListByDataflowProfile_MaximumSet.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void dataflowGraphListByDataflowProfileMaximumSet(
        com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflowGraphs().listByDataflowProfile("rgiotoperations", "resource-123", "resource-123",
            com.azure.core.util.Context.NONE);
    }
}
