/** Samples for IntegrationAccounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccounts_ListBySubscription.json
     */
    /**
     * Sample code: List integration accounts by subscription.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void listIntegrationAccountsBySubscription(com.azure.resourcemanager.logic.LogicManager manager) {
        manager.integrationAccounts().list(null, com.azure.core.util.Context.NONE);
    }
}
