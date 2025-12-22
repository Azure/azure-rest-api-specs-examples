
/**
 * Samples for IntelligencePacks Disable.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/WorkspacesDisableIntelligencePack.json
     */
    /**
     * Sample code: IntelligencePacksDisable.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void intelligencePacksDisable(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.intelligencePacks().disableWithResponse("rg1", "TestLinkWS", "ChangeTracking",
            com.azure.core.util.Context.NONE);
    }
}
