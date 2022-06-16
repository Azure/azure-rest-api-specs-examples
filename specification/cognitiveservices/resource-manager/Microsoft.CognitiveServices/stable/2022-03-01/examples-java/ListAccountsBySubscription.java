import com.azure.core.util.Context;

/** Samples for Accounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/ListAccountsBySubscription.json
     */
    /**
     * Sample code: List Accounts by Subscription.
     *
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listAccountsBySubscription(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.accounts().list(Context.NONE);
    }
}
