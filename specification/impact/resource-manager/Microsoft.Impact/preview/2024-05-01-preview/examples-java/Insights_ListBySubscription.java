
/**
 * Samples for Insights ListBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/Insights_ListBySubscription.json
     */
    /**
     * Sample code: List Insight resources by workloadImpactName.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void listInsightResourcesByWorkloadImpactName(
        com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        manager.insights().listBySubscription("impactid22", com.azure.core.util.Context.NONE);
    }
}
