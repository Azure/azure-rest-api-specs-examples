
/**
 * Samples for WebApps ListBasicPublishingCredentialsPolicies.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListPublishingCredentialsPolicies.json
     */
    /**
     * Sample code: List Publishing Credentials Policies.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        listPublishingCredentialsPolicies(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().listBasicPublishingCredentialsPolicies("testrg123", "testsite",
            com.azure.core.util.Context.NONE);
    }
}
