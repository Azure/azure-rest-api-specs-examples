
/**
 * Samples for DrillRuns GenerateReport.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/DrillRuns_GenerateReport_MaximumSet_Gen.json
     */
    /**
     * Sample code: DrillRuns_GenerateReport_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void drillRunsGenerateReportMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.drillRuns().generateReport("sampleServiceGroupName", "qmn", "drill1",
            "ca92602e-53bf-43d2-ae62-d3fc940474b3", com.azure.core.util.Context.NONE);
    }
}
