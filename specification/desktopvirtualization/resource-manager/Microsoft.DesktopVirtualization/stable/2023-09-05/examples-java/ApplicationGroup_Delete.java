/** Samples for ApplicationGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/ApplicationGroup_Delete.json
     */
    /**
     * Sample code: ApplicationGroup_Delete.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void applicationGroupDelete(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .applicationGroups()
            .deleteByResourceGroupWithResponse("resourceGroup1", "applicationGroup1", com.azure.core.util.Context.NONE);
    }
}
