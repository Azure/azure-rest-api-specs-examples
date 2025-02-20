
/**
 * Samples for Insights Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/Insights_Get_diagnostics.json
     */
    /**
     * Sample code: Get Insight sample for Diagnostics category.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void getInsightSampleForDiagnosticsCategory(
        com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        manager.insights().getWithResponse("impactid", "insight1", com.azure.core.util.Context.NONE);
    }
}
