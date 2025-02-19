
/**
 * Samples for Insights Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/Insights_Get_mitigationAction.json
     */
    /**
     * Sample code: Get Insight sample for MitigationAction category.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void getInsightSampleForMitigationActionCategory(
        com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        manager.insights().getWithResponse("impactId", "HPCUASucceeded", com.azure.core.util.Context.NONE);
    }
}
