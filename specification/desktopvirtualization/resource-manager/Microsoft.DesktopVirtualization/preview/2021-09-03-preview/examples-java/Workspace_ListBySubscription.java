import com.azure.core.util.Context;

/** Samples for Workspaces List. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/Workspace_ListBySubscription.json
     */
    /**
     * Sample code: Workspace_ListBySubscription.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void workspaceListBySubscription(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.workspaces().list(Context.NONE);
    }
}
