import com.azure.core.util.Context;

/** Samples for ApplicationGroups Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/ApplicationGroup_Delete.json
     */
    /**
     * Sample code: ApplicationGroup_Delete.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void applicationGroupDelete(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.applicationGroups().deleteWithResponse("resourceGroup1", "applicationGroup1", Context.NONE);
    }
}
