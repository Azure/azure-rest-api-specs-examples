
import com.azure.resourcemanager.cdn.models.ValidateCustomDomainInput;

/**
 * Samples for Endpoints ValidateCustomDomain.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Endpoints_ValidateCustomDomain.json
     */
    /**
     * Sample code: Endpoints_ValidateCustomDomain.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void endpointsValidateCustomDomain(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getEndpoints().validateCustomDomainWithResponse("RG", "profile1", "endpoint1",
            new ValidateCustomDomainInput().withHostname("www.someDomain.com"), com.azure.core.util.Context.NONE);
    }
}
