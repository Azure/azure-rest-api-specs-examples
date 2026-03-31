
/**
 * Samples for WebApps UpdateMachineKey.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/UpdateMachineKey.json
     */
    /**
     * Sample code: Updates the machine key for a site.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void updatesTheMachineKeyForASite(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().updateMachineKeyWithResponse("rg", "contoso",
            com.azure.core.util.Context.NONE);
    }
}
