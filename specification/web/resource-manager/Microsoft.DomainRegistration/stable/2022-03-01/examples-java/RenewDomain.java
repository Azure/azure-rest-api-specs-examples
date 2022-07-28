import com.azure.core.util.Context;

/** Samples for Domains Renew. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2022-03-01/examples/RenewDomain.json
     */
    /**
     * Sample code: Renew an existing domain.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void renewAnExistingDomain(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDomains().renewWithResponse("RG", "example.com", Context.NONE);
    }
}
