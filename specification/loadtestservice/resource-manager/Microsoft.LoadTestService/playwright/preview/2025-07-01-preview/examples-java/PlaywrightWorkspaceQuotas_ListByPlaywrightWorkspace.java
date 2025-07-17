
/**
 * Samples for PlaywrightWorkspaceQuotas ListByPlaywrightWorkspace.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01-preview/PlaywrightWorkspaceQuotas_ListByPlaywrightWorkspace.json
     */
    /**
     * Sample code: PlaywrightWorkspaceQuotas_ListByPlaywrightWorkspace.
     * 
     * @param manager Entry point to PlaywrightManager.
     */
    public static void playwrightWorkspaceQuotasListByPlaywrightWorkspace(
        com.azure.resourcemanager.playwright.PlaywrightManager manager) {
        manager.playwrightWorkspaceQuotas().listByPlaywrightWorkspace("dummyrg", "myWorkspace",
            com.azure.core.util.Context.NONE);
    }
}
