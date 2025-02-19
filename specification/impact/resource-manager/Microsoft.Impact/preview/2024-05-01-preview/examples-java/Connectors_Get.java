
/**
 * Samples for Connectors Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/Connectors_Get.json
     */
    /**
     * Sample code: Connectors_Get.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void connectorsGet(com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        manager.connectors().getWithResponse("testconnector1", com.azure.core.util.Context.NONE);
    }
}
