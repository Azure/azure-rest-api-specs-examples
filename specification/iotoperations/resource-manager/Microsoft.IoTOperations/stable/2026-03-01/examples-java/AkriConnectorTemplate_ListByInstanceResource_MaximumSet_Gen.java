
/**
 * Samples for AkriConnectorTemplate ListByInstanceResource.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/AkriConnectorTemplate_ListByInstanceResource_MaximumSet_Gen.json
     */
    /**
     * Sample code: AkriConnectorTemplate_ListByInstanceResource_MaximumSet.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void akriConnectorTemplateListByInstanceResourceMaximumSet(
        com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.akriConnectorTemplates().listByInstanceResource("rgiotoperations", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
