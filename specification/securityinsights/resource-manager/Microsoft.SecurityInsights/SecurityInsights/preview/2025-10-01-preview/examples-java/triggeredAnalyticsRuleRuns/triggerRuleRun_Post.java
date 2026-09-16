
import com.azure.resourcemanager.securityinsights.models.AnalyticsRuleRunTrigger;
import java.time.OffsetDateTime;

/**
 * Samples for AlertRuleOperation TriggerRuleRun.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/triggeredAnalyticsRuleRuns/triggerRuleRun_Post.json
     */
    /**
     * Sample code: triggerRuleRun_Post.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void triggerRuleRunPost(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.alertRuleOperations().triggerRuleRun("myRg", "myWorkspace", "65360bb0-8986-4ade-a89d-af3cf44d28aa",
            new AnalyticsRuleRunTrigger().withExecutionTimeUtc(OffsetDateTime.parse("2022-12-22T15:37:03.074Z")),
            com.azure.core.util.Context.NONE);
    }
}
