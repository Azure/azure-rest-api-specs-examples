/** Samples for IntegrationServiceEnvironments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_Delete.json
     */
    /**
     * Sample code: Delete an integration account.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void deleteAnIntegrationAccount(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationServiceEnvironments()
            .deleteByResourceGroupWithResponse(
                "testResourceGroup", "testIntegrationServiceEnvironment", com.azure.core.util.Context.NONE);
    }
}
