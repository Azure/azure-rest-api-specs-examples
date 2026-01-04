
/**
 * Samples for WebApps UpdateMachineKey.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/UpdateMachineKey.json
     */
    /**
     * Sample code: Updates the machine key for a site.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updatesTheMachineKeyForASite(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().updateMachineKeyWithResponse("rg", "contoso",
            com.azure.core.util.Context.NONE);
    }
}
