
/**
 * Samples for IntelligencePacks Enable.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/OperationalInsights/stable/2025-
     * 07-01/examples/WorkspacesEnableIntelligencePack.json
     */
    /**
     * Sample code: IntelligencePacksEnable.
     * 
     * @param manager Entry point to LogAnalyticsManager.
     */
    public static void intelligencePacksEnable(com.azure.resourcemanager.loganalytics.LogAnalyticsManager manager) {
        manager.intelligencePacks().enableWithResponse("rg1", "TestLinkWS", "ChangeTracking",
            com.azure.core.util.Context.NONE);
    }
}
