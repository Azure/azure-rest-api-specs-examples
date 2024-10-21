
/**
 * Samples for Dataflow Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-15-preview/Dataflow_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Dataflow_Delete.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void dataflowDelete(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.dataflows().delete("rgiotoperations", "resource-name123", "resource-name123", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
