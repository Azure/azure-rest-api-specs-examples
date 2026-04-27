
/**
 * Samples for PlaywrightWorkspaces GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01-preview/PlaywrightWorkspaces_Get.json
     */
    /**
     * Sample code: PlaywrightWorkspaces_Get.
     * 
     * @param manager Entry point to PlaywrightManager.
     */
    public static void playwrightWorkspacesGet(com.azure.resourcemanager.playwright.PlaywrightManager manager) {
        manager.playwrightWorkspaces().getByResourceGroupWithResponse("dummyrg", "myWorkspace",
            com.azure.core.util.Context.NONE);
    }
}
