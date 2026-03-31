
/**
 * Samples for WebApps ListBasicPublishingCredentialsPoliciesSlot.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListPublishingCredentialsPoliciesSlot.json
     */
    /**
     * Sample code: List Publishing Credentials Policies.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void
        listPublishingCredentialsPolicies(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().listBasicPublishingCredentialsPoliciesSlot("testrg123", "testsite",
            "staging", com.azure.core.util.Context.NONE);
    }
}
