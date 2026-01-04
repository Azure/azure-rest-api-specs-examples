
/**
 * Samples for Domains List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/domainregistration/resource-manager/Microsoft.DomainRegistration/DomainRegistration/stable/2024-11-
     * 01/examples/ListDomainsBySubscription.json
     */
    /**
     * Sample code: List domains by subscription.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDomainsBySubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDomains().list(com.azure.core.util.Context.NONE);
    }
}
