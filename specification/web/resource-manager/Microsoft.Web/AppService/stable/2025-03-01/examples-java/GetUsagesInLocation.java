
/**
 * Samples for GetUsagesInLocation List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetUsagesInLocation.json
     */
    /**
     * Sample code: Get usages in location for subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getUsagesInLocationForSubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getGetUsagesInLocations().list("West US",
            com.azure.core.util.Context.NONE);
    }
}
