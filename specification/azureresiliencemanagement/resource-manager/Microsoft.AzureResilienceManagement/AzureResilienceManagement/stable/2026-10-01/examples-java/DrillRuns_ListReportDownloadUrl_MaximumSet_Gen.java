
import com.azure.resourcemanager.resiliencemanagement.models.DrillReportFormat;
import com.azure.resourcemanager.resiliencemanagement.models.ListReportDownloadUrlRequest;

/**
 * Samples for DrillRuns ListReportDownloadUrl.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/DrillRuns_ListReportDownloadUrl_MaximumSet_Gen.json
     */
    /**
     * Sample code: DrillRuns_ListReportDownloadUrl_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void drillRunsListReportDownloadUrlMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.drillRuns().listReportDownloadUrl("sampleServiceGroupName", "3f2b1c4d-5e6f-4a7b-8c9d-0e1f2a3b4c5d",
            "drill1", "ca92602e-53bf-43d2-ae62-d3fc940474b3",
            new ListReportDownloadUrlRequest().withFormat(DrillReportFormat.HTML), com.azure.core.util.Context.NONE);
    }
}
