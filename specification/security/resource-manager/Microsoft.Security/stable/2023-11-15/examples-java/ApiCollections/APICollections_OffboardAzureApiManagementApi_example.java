
/**
 * Samples for ApiCollections OffboardAzureApiManagementApi.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2023-11-15/examples/ApiCollections/
     * APICollections_OffboardAzureApiManagementApi_example.json
     */
    /**
     * Sample code: Offboard an Azure API Management API from Microsoft Defender for APIs.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void offboardAnAzureAPIManagementAPIFromMicrosoftDefenderForAPIs(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.apiCollections().offboardAzureApiManagementApiWithResponse("rg1", "apimService1", "echo-api",
            com.azure.core.util.Context.NONE);
    }
}
