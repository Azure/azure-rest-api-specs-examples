import com.azure.core.util.Context;

/** Samples for ApplicationGroups ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/ApplicationGroup_ListByResourceGroup.json
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
            .listByResourceGroup("resourceGroup1", "applicationGroupType eq 'RailApplication'", Context.NONE);
    }
}
