
/**
 * Samples for AkriConnectorTemplate Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/AkriConnectorTemplate_Get_Managed_Rest.json
     */
    /**
     * Sample code: AkriConnectorTemplate_Get_Managed_Rest.
     * 
     * @param manager Entry point to IoTOperationsManager.
     */
    public static void
        akriConnectorTemplateGetManagedRest(com.azure.resourcemanager.iotoperations.IoTOperationsManager manager) {
        manager.akriConnectorTemplates().getWithResponse("rgiotoperations", "resource-name123", "resource-name123",
            com.azure.core.util.Context.NONE);
    }
}
