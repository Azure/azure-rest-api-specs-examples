
/**
 * Samples for ApiCollections ListByAzureApiManagementService.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2023-11-15/examples/ApiCollections/
     * APICollections_ListByAzureApiManagementService_example.json
     */
    /**
     * Sample code: Gets a list of Azure API Management APIs that have been onboarded to Microsoft Defender for APIs.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getsAListOfAzureAPIManagementAPIsThatHaveBeenOnboardedToMicrosoftDefenderForAPIs(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.apiCollections().listByAzureApiManagementService("rg1", "apimService1",
            com.azure.core.util.Context.NONE);
    }
}
