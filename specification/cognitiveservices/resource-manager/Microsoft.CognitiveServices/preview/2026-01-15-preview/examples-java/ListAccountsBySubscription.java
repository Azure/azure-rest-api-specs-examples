
/**
 * Samples for Accounts List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-15-preview/ListAccountsBySubscription.json
     */
    /**
     * Sample code: List Accounts by Subscription.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listAccountsBySubscription(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().list(com.azure.core.util.Context.NONE);
    }
}
