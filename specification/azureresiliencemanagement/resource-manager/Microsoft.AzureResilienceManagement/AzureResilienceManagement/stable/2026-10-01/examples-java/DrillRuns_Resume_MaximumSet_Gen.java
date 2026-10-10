
/**
 * Samples for DrillRuns Resume.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/DrillRuns_Resume_MaximumSet_Gen.json
     */
    /**
     * Sample code: DrillRuns_Resume_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void
        drillRunsResumeMaximumSet(com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.drillRuns().resume("sampleServiceGroupName", "qmn", "drill1", "ca92602e-53bf-43d2-ae62-d3fc940474b3",
            com.azure.core.util.Context.NONE);
    }
}
