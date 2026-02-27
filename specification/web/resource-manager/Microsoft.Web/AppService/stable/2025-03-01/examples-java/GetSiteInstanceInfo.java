
/**
 * Samples for WebApps GetInstanceInfoSlot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetSiteInstanceInfo.json
     */
    /**
     * Sample code: Get site instance info.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getSiteInstanceInfo(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().getInstanceInfoSlotWithResponse("testrg123", "tests346",
            "134987120", "staging", com.azure.core.util.Context.NONE);
    }
}
