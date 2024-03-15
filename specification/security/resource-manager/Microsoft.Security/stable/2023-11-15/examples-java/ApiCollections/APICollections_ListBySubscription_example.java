
/**
 * Samples for ApiCollections List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2023-11-15/examples/ApiCollections/
     * APICollections_ListBySubscription_example.json
     */
    /**
     * Sample code: Gets a list of API collections within a subscription that have been onboarded to Microsoft Defender
     * for APIs.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getsAListOfAPICollectionsWithinASubscriptionThatHaveBeenOnboardedToMicrosoftDefenderForAPIs(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.apiCollections().list(com.azure.core.util.Context.NONE);
    }
}
