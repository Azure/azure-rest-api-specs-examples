
import com.azure.resourcemanager.resiliencemanagement.models.DrillRunAddNotesRequest;

/**
 * Samples for DrillRuns AddNotes.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/DrillRuns_AddNotes_MaximumSet_Gen.json
     */
    /**
     * Sample code: DrillRuns_AddNotes_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void drillRunsAddNotesMaximumSet(
        com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.drillRuns().addNotes("sampleServiceGroupName", "qmn", "drill1", "ca92602e-53bf-43d2-ae62-d3fc940474b3",
            new DrillRunAddNotesRequest().withNotes("wubqjajveatmwcglo"), com.azure.core.util.Context.NONE);
    }
}
