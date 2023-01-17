/** Samples for IntegrationAccounts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccounts_Delete.json
     */
    /**
     * Sample code: Delete an integration account.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void deleteAnIntegrationAccount(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccounts()
            .deleteByResourceGroupWithResponse(
                "testResourceGroup", "testIntegrationAccount", com.azure.core.util.Context.NONE);
    }
}
