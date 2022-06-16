import com.azure.core.util.Context;
import com.azure.resourcemanager.desktopvirtualization.models.DesktopPatch;

/** Samples for Desktops Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/preview/2021-09-03-preview/examples/Desktop_Update.json
     */
    /**
     * Sample code: Desktop_Update.
     *
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void desktopUpdate(
        com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager
            .desktops()
            .updateWithResponse(
                "resourceGroup1",
                "applicationGroup1",
                "SessionDesktop",
                new DesktopPatch().withDescription("des1").withFriendlyName("friendly"),
                Context.NONE);
    }
}
