/** Samples for Workspaces Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2022-09-09/examples/Workspace_Delete.json
     */
    /**
     * Sample code: Workspace_Delete.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void workspaceDelete(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .workspaces()
            .deleteByResourceGroupWithResponse("resourceGroup1", "workspace1", com.azure.core.util.Context.NONE);
    }
}
