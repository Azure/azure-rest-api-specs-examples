import com.azure.core.util.Context;

/** Samples for ApplicationGroups List. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/ApplicationGroup_ListBySubscription.json
     */
    /**
     * Sample code: ApplicationGroup_List.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void applicationGroupList(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.applicationGroups().list("applicationGroupType eq 'RailApplication'", Context.NONE);
    }
}
