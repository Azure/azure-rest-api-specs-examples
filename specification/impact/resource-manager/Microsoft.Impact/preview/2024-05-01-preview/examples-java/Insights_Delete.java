
/**
 * Samples for Insights Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/Insights_Delete.json
     */
    /**
     * Sample code: Delete an Insight.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void deleteAnInsight(com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        manager.insights().deleteByResourceGroupWithResponse("impactid22", "insightId12",
            com.azure.core.util.Context.NONE);
    }
}
