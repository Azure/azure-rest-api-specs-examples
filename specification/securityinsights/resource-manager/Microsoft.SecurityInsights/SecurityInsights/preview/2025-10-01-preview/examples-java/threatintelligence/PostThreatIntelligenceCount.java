
import com.azure.resourcemanager.securityinsights.models.CountQuery;
import com.azure.resourcemanager.securityinsights.models.TiType;

/**
 * Samples for ThreatIntelligence Count.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/threatintelligence/PostThreatIntelligenceCount.json
     */
    /**
     * Sample code: Get TI object count.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getTIObjectCount(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.threatIntelligences().countWithResponse("myRg", "myWorkspace", TiType.MAIN, new CountQuery(),
            com.azure.core.util.Context.NONE);
    }
}
