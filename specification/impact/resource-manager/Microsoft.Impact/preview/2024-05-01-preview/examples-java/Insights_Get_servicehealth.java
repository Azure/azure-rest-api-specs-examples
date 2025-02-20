
/**
 * Samples for Insights Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-05-01-preview/Insights_Get_servicehealth.json
     */
    /**
     * Sample code: Get Insight sample for service health category.
     * 
     * @param manager Entry point to ImpactReportingManager.
     */
    public static void getInsightSampleForServiceHealthCategory(
        com.azure.resourcemanager.impactreporting.ImpactReportingManager manager) {
        manager.insights().getWithResponse("impactid", "insightname", com.azure.core.util.Context.NONE);
    }
}
