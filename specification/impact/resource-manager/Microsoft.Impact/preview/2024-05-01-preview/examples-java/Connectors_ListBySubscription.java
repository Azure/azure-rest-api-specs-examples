
/**
 * Samples for Connectors List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/Connectors_ListBySubscription.json
     */
    /**
     * Sample code: Connectors_ListBySubscription.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void
        connectorsListBySubscription(com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        manager.connectors().list(com.azure.core.util.Context.NONE);
    }
}
