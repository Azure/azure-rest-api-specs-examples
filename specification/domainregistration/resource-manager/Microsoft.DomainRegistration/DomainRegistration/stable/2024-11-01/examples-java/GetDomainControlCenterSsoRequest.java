
/**
 * Samples for Domains GetControlCenterSsoRequest.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/domainregistration/resource-manager/Microsoft.DomainRegistration/DomainRegistration/stable/2024-11-
     * 01/examples/GetDomainControlCenterSsoRequest.json
     */
    /**
     * Sample code: Get Domain Control Center Sso Request.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDomainControlCenterSsoRequest(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDomains()
            .getControlCenterSsoRequestWithResponse(com.azure.core.util.Context.NONE);
    }
}
