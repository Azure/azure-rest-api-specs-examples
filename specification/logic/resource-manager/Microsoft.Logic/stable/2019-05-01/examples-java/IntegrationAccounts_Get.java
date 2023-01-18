/** Samples for IntegrationAccounts GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccounts_Get.json
     */
    /**
     * Sample code: Get integration account by name.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getIntegrationAccountByName(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccounts()
            .getByResourceGroupWithResponse(
                "testResourceGroup", "testIntegrationAccount", com.azure.core.util.Context.NONE);
    }
}
