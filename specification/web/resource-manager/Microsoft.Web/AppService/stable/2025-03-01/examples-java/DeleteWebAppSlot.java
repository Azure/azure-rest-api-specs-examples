
/**
 * Samples for WebApps DeleteSlot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/DeleteWebAppSlot.json
     */
    /**
     * Sample code: Delete Web App Slot.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteWebAppSlot(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().deleteSlotWithResponse("testrg123", "sitef6141",
            "staging", null, null, com.azure.core.util.Context.NONE);
    }
}
