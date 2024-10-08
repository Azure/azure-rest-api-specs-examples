
/**
 * Samples for SessionHosts Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/desktopvirtualization/resource-manager/Microsoft.DesktopVirtualization/stable/2024-04-03/examples/
     * SessionHost_Get.json
     */
    /**
     * Sample code: SessionHost_Get.
     * 
     * @param manager Entry point to DesktopVirtualizationManager.
     */
    public static void
        sessionHostGet(com.azure.resourcemanager.desktopvirtualization.DesktopVirtualizationManager manager) {
        manager.sessionHosts().getWithResponse("resourceGroup1", "hostPool1", "sessionHost1.microsoft.com",
            com.azure.core.util.Context.NONE);
    }
}
