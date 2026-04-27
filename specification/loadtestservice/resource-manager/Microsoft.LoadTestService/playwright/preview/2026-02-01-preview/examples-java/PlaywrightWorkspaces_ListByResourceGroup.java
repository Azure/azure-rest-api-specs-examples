
/**
 * Samples for PlaywrightWorkspaces ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01-preview/PlaywrightWorkspaces_ListByResourceGroup.json
     */
    /**
     * Sample code: PlaywrightWorkspaces_ListByResourceGroup.
     * 
     * @param manager Entry point to PlaywrightManager.
     */
    public static void
        playwrightWorkspacesListByResourceGroup(com.azure.resourcemanager.playwright.PlaywrightManager manager) {
        manager.playwrightWorkspaces().listByResourceGroup("dummyrg", com.azure.core.util.Context.NONE);
    }
}
