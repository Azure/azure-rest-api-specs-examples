import com.azure.core.util.Context;

/** Samples for Domains List. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2022-03-01/examples/ListDomainsBySubscription.json
     */
    /**
     * Sample code: List domains by subscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDomainsBySubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDomains().list(Context.NONE);
    }
}
