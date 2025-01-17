
import com.azure.resourcemanager.securityinsights.models.ThreatIntelligenceAppendTags;
import java.util.Arrays;

/**
 * Samples for ThreatIntelligenceIndicator AppendTags.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/securityinsights/resource-manager/Microsoft.SecurityInsights/stable/2022-11-01/examples/
     * threatintelligence/AppendTagsThreatIntelligence.json
     */
    /**
     * Sample code: Append tags to a threat intelligence indicator.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void appendTagsToAThreatIntelligenceIndicator(
        com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.threatIntelligenceIndicators().appendTagsWithResponse("myRg", "myWorkspace",
            "d9cd6f0b-96b9-3984-17cd-a779d1e15a93",
            new ThreatIntelligenceAppendTags().withThreatIntelligenceTags(Arrays.asList("tag1", "tag2")),
            com.azure.core.util.Context.NONE);
    }
}
