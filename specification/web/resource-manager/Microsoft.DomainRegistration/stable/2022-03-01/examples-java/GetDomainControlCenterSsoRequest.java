import com.azure.core.util.Context;

/** Samples for Domains GetControlCenterSsoRequest. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.DomainRegistration/stable/2022-03-01/examples/GetDomainControlCenterSsoRequest.json
     */
    /**
     * Sample code: Get Domain Control Center Sso Request.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDomainControlCenterSsoRequest(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDomains().getControlCenterSsoRequestWithResponse(Context.NONE);
    }
}
