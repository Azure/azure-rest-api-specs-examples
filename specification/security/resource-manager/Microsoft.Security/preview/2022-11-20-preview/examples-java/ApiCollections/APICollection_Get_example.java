import com.azure.core.util.Context;

/** Samples for ApiCollection Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-11-20-preview/examples/ApiCollections/APICollection_Get_example.json
     */
    /**
     * Sample code: Gets an Azure API Management API if it has been onboarded to Defender for APIs.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getsAnAzureAPIManagementAPIIfItHasBeenOnboardedToDefenderForAPIs(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.apiCollections().getWithResponse("rg1", "apimService1", "echo-api", Context.NONE);
    }
}
