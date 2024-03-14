
/**
 * Samples for ApiCollections OnboardAzureApiManagementApi.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2023-11-15/examples/ApiCollections/
     * APICollections_OnboardAzureApiManagementApi_example.json
     */
    /**
     * Sample code: Onboard an Azure API Management API to Microsoft Defender for APIs.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void onboardAnAzureAPIManagementAPIToMicrosoftDefenderForAPIs(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.apiCollections().onboardAzureApiManagementApi("rg1", "apimService1", "echo-api",
            com.azure.core.util.Context.NONE);
    }
}
