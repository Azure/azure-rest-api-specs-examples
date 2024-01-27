
/**
 * Samples for DevToolPortals List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2023-12-01/examples/DevToolPortals_List.
     * json
     */
    /**
     * Sample code: DevToolPortals_List.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void devToolPortalsList(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.springServices().manager().serviceClient().getDevToolPortals().list("myResourceGroup", "myservice",
            com.azure.core.util.Context.NONE);
    }
}
