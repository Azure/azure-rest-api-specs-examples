
/**
 * Samples for PlaywrightWorkspaces List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01-preview/PlaywrightWorkspaces_ListBySubscription.json
     */
    /**
     * Sample code: PlaywrightWorkspaces_ListBySubscription.
     * 
     * @param manager Entry point to PlaywrightManager.
     */
    public static void
        playwrightWorkspacesListBySubscription(com.azure.resourcemanager.playwright.PlaywrightManager manager) {
        manager.playwrightWorkspaces().list(com.azure.core.util.Context.NONE);
    }
}
