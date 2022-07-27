import com.azure.core.util.Context;

/** Samples for Domains Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2022-03-01/examples/DeleteAppServiceDomain.json
     */
    /**
     * Sample code: Delete App Service Domain.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deleteAppServiceDomain(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getDomains()
            .deleteWithResponse("testrg123", "example.com", true, Context.NONE);
    }
}
