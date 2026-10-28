
import com.azure.resourcemanager.cdn.models.ValidateCustomDomainInput;

/**
 * Samples for AfdEndpoints ValidateCustomDomain.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDEndpoints_ValidateCustomDomain.json
     */
    /**
     * Sample code: Endpoints_ValidateCustomDomain.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void endpointsValidateCustomDomain(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdEndpoints().validateCustomDomainWithResponse("RG", "profile1", "endpoint1",
            new ValidateCustomDomainInput().withHostname("www.someDomain.com"), com.azure.core.util.Context.NONE);
    }
}
