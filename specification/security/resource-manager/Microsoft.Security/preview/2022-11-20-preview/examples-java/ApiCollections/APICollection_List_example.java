/** Samples for ApiCollection List. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-11-20-preview/examples/ApiCollections/APICollection_List_example.json
     */
    /**
     * Sample code: Gets a list of Azure API Management APIs that have been onboarded to Defender for APIs.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getsAListOfAzureAPIManagementAPIsThatHaveBeenOnboardedToDefenderForAPIs(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.apiCollections().list("rg1", "apimService1", com.azure.core.util.Context.NONE);
    }
}
