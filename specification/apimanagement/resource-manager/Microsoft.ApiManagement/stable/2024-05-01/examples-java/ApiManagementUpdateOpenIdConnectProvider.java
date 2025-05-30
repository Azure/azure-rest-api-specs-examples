
import com.azure.resourcemanager.apimanagement.models.OpenidConnectProviderContract;

/**
 * Samples for OpenIdConnectProvider Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementUpdateOpenIdConnectProvider.json
     */
    /**
     * Sample code: ApiManagementUpdateOpenIdConnectProvider.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementUpdateOpenIdConnectProvider(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        OpenidConnectProviderContract resource = manager.openIdConnectProviders()
            .getWithResponse("rg1", "apimService1", "templateOpenIdConnect2", com.azure.core.util.Context.NONE)
            .getValue();
        resource.update().withClientSecret("updatedsecret").withUseInTestConsole(false).withUseInApiDocumentation(true)
            .withIfMatch("*").apply();
    }
}
