import com.azure.core.util.Context;

/** Samples for Domains GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2022-03-01/examples/GetDomain.json
     */
    /**
     * Sample code: Get Domain.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDomain(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getDomains()
            .getByResourceGroupWithResponse("testrg123", "example.com", Context.NONE);
    }
}
