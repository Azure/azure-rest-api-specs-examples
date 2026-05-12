
/**
 * Samples for AkriConnectorTemplate Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/AkriConnectorTemplate_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: AkriConnectorTemplate_Delete_MaximumSet.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        akriConnectorTemplateDeleteMaximumSet(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.akriConnectorTemplates().delete("rgiotoperations", "resource-name123", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
