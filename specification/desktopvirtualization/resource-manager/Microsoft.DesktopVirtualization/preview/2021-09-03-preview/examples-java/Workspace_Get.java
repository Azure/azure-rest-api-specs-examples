import com.azure.core.util.Context;

/** Samples for Workspaces GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/Workspace_Get.json
     */
    /**
     * Sample code: Workspace_Get.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void workspaceGet(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.workspaces().getByResourceGroupWithResponse("resourceGroup1", "workspace1", Context.NONE);
    }
}
