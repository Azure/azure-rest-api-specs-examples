
/**
 * Samples for AkriConnector ListByTemplate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/AkriConnector_ListByTemplate_MaximumSet_Gen.json
     */
    /**
     * Sample code: AkriConnector_ListByTemplate_MaximumSet.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        akriConnectorListByTemplateMaximumSet(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.akriConnectors().listByTemplate("rgiotoperations", "resource-name123", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
