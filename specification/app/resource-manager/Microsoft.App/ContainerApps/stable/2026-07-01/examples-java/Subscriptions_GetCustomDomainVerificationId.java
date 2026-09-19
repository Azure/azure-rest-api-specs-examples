
/**
 * Samples for ResourceProvider GetCustomDomainVerificationId.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Subscriptions_GetCustomDomainVerificationId.json
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
