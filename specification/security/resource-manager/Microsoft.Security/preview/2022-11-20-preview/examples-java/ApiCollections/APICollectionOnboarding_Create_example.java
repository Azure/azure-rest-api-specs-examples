/** Samples for ApiCollectionOnboarding Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2022-11-20-preview/examples/ApiCollections/APICollectionOnboarding_Create_example.json
     */
    /**
     * Sample code: Onboard an Azure API Management API to Defender for APIs.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void onboardAnAzureAPIManagementAPIToDefenderForAPIs(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .apiCollectionOnboardings()
            .createWithResponse("rg1", "apimService1", "echo-api", com.azure.core.util.Context.NONE);
    }
}
