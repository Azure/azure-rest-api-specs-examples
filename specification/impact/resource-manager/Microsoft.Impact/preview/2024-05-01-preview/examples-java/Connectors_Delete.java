
/**
 * Samples for Connectors Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/Connectors_Delete.json
     */
    /**
     * Sample code: Connectors_Delete.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void connectorsDelete(com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        manager.connectors().deleteWithResponse("testconnector1", com.azure.core.util.Context.NONE);
    }
}
