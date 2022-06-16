import com.azure.core.util.Context;
import com.azure.resourcemanager.cdn.models.ValidateCustomDomainInput;

/** Samples for Endpoints ValidateCustomDomain. */
public final class Main {
    /*
     * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/Endpoints_ValidateCustomDomain.json
     */
    /**
     * Sample code: Endpoints_ValidateCustomDomain.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void endpointsValidateCustomDomain(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .cdnProfiles()
            .manager()
            .serviceClient()
            .getEndpoints()
            .validateCustomDomainWithResponse(
                "RG",
                "profile1",
                "endpoint1",
                new ValidateCustomDomainInput().withHostname("www.someDomain.com"),
                Context.NONE);
    }
}
