/** Samples for ApplicationGroups ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2023-09-05/examples/ApplicationGroup_ListByResourceGroup.json
     */
    /**
     * Sample code: ApplicationGroup_ListByResourceGroup.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void applicationGroupListByResourceGroup(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .applicationGroups()
            .listByResourceGroup(
                "resourceGroup1",
                "applicationGroupType eq 'RailApplication'",
                10,
                true,
                0,
                com.azure.core.util.Context.NONE);
    }
}
