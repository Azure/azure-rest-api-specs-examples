/** Samples for Accounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/graphservicesprod/resource-manager/Microsoft.GraphServices/stable/2023-04-13/examples/Accounts_List_Sub.json
     */
    /**
     * Sample code: Get list of accounts by subscription.
     *
     * @param manager Entry point to GraphServicesManager.
     */
    public static void getListOfAccountsBySubscription(
        com.azure.resourcemanager.graphservices.GraphServicesManager manager) {
        manager.accounts().list(com.azure.core.util.Context.NONE);
    }
}
