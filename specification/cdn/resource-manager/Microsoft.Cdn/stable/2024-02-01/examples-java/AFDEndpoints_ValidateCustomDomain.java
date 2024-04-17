
import com.azure.resourcemanager.cdn.models.ValidateCustomDomainInput;

/**
 * Samples for AfdEndpoints ValidateCustomDomain.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/AFDEndpoints_ValidateCustomDomain.
     * json
     */
    /**
     * Sample code: Endpoints_ValidateCustomDomain.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointsValidateCustomDomain(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.cdnProfiles().manager().serviceClient().getAfdEndpoints().validateCustomDomainWithResponse("RG",
            "profile1", "endpoint1", new ValidateCustomDomainInput().withHostname("www.someDomain.com"),
            com.azure.core.util.Context.NONE);
    }
}
