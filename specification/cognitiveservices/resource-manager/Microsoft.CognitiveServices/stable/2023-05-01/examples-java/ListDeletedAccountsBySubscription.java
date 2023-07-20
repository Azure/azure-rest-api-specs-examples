/** Samples for DeletedAccounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2023-05-01/examples/ListDeletedAccountsBySubscription.json
     */
    /**
     * Sample code: List Deleted Accounts by Subscription.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listDeletedAccountsBySubscription(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.deletedAccounts().list(com.azure.core.util.Context.NONE);
    }
}
