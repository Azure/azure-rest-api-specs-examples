
/**
 * Samples for Metadata Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01-preview/metadata/PutMetadataMinimal.json
     */
    /**
     * Sample code: Create/update minimal metadata.
     * 
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void
        createUpdateMinimalMetadata(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.metadatas().define("metadataName").withExistingWorkspace("myRg", "myWorkspace")
            .withContentId("c00ee137-7475-47c8-9cce-ec6f0f1bedd0")
            .withParentId(
                "/subscriptions/2e1dc338-d04d-4443-b721-037eff4fdcac/resourceGroups/myRg/providers/Microsoft.OperationalInsights/workspaces/myWorkspace/providers/Microsoft.SecurityInsights/alertRules/ruleName")
            .withKind("AnalyticsRule").create();
    }
}
