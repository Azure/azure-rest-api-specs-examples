
/**
 * Samples for AppServiceEnvironments List.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/
     * AppServiceEnvironments_List.json
     */
    /**
     * Sample code: Get all App Service Environments for a subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void
        getAllAppServiceEnvironmentsForASubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getAppServiceEnvironments().list(com.azure.core.util.Context.NONE);
    }
}
