
/**
 * Samples for ResourceProvider GetCustomDomainVerificationId.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/app/resource-manager/Microsoft.App/ContainerApps/preview/2025-02-02-preview/examples/
     * Subscriptions_GetCustomDomainVerificationId.json
     */
    /**
     * Sample code: List all operations.
     * 
     * @param manager Entry point to ContainerAppsApiManager.
     */
    public static void listAllOperations(com.azure.resourcemanager.appcontainers.ContainerAppsApiManager manager) {
        manager.resourceProviders().getCustomDomainVerificationIdWithResponse(com.azure.core.util.Context.NONE);
    }
}
