import com.azure.core.util.Context;

/** Samples for Workspaces ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/Workspace_ListByResourceGroup.json
     */
    /**
     * Sample code: Workspace_ListByResourceGroup.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void workspaceListByResourceGroup(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.workspaces().listByResourceGroup("resourceGroup1", Context.NONE);
    }
}
