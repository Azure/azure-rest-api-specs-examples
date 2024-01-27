
/**
 * Samples for TopLevelDomains Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.DomainRegistration/stable/2023-01-01/examples/GetTopLevelDomain.json
     */
    /**
     * Sample code: Get Top Level Domain.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getTopLevelDomain(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getTopLevelDomains().getWithResponse("com",
            com.azure.core.util.Context.NONE);
    }
}
