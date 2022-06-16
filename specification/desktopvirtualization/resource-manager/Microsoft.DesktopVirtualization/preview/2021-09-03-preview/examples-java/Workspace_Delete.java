import com.azure.core.util.Context;

/** Samples for Workspaces Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/Workspace_Delete.json
     */
    /**
     * Sample code: Workspace_Delete.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void workspaceDelete(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.workspaces().deleteWithResponse("resourceGroup1", "workspace1", Context.NONE);
    }
}
