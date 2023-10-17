/** Samples for Workspaces GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/Workspace_Get.json
     */
    /**
     * Sample code: Workspace_Get.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void workspaceGet(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .workspaces()
            .getByResourceGroupWithResponse("resourceGroup1", "workspace1", com.azure.core.util.Context.NONE);
    }
}
