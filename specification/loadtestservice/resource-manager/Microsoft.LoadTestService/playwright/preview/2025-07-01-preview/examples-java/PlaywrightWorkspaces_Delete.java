
/**
 * Samples for PlaywrightWorkspaces Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-01-preview/PlaywrightWorkspaces_Delete.json
     */
    /**
     * Sample code: PlaywrightWorkspaces_Delete.
     * 
     * @param manager Entry point to PlaywrightManager.
     */
    public static void playwrightWorkspacesDelete(com.azure.resourcemanager.playwright.PlaywrightManager manager) {
        manager.playwrightWorkspaces().delete("dummyrg", "myWorkspace", com.azure.core.util.Context.NONE);
    }
}
