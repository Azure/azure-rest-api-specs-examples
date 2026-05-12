
/**
 * Samples for AkriConnector Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/AkriConnector_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: AkriConnector_Delete_MaximumSet.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        akriConnectorDeleteMaximumSet(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.akriConnectors().delete("rgiotoperations", "resource-name123", "resource-name123", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
