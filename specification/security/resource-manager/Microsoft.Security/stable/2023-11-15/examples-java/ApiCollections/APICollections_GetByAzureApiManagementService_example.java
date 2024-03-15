
/**
 * Samples for ApiCollections GetByAzureApiManagementService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2023-11-15/examples/ApiCollections/
     * APICollections_GetByAzureApiManagementService_example.json
     */
    /**
     * Sample code: Gets an Azure API Management API if it has been onboarded to Microsoft Defender for APIs.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getsAnAzureAPIManagementAPIIfItHasBeenOnboardedToMicrosoftDefenderForAPIs(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.apiCollections().getByAzureApiManagementServiceWithResponse("rg1", "apimService1", "echo-api",
            com.azure.core.util.Context.NONE);
    }
}
