
/**
 * Samples for ResourceProvider GetCustomDomainVerificationId.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2024-03-01/examples/
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
