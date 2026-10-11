
import com.azure.resourcemanager.resiliencemanagement.models.DrillMode;
import com.azure.resourcemanager.resiliencemanagement.models.DrillStartRequest;

/**
 * Samples for Drills Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Drills_Start_MaximumSet_Gen.json
     */
    /**
     * Sample code: Drills_Start_MaximumSet.
     * 
     * @param manager Entry point to ResilienceManagementManager.
     */
    public static void
        drillsStartMaximumSet(com.azure.resourcemanager.resiliencemanagement.ResilienceManagementManager manager) {
        manager.drills().start("sampleServiceGroupName", "qmn", "drill1",
            new DrillStartRequest().withMode(DrillMode.FAILOVER), com.azure.core.util.Context.NONE);
    }
}
