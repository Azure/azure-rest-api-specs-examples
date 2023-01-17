/** Samples for IntegrationServiceEnvironments Restart. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_Restart.json
     */
    /**
     * Sample code: Restarts an integration service environment.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void restartsAnIntegrationServiceEnvironment(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationServiceEnvironments()
            .restartWithResponse(
                "testResourceGroup", "testIntegrationServiceEnvironment", com.azure.core.util.Context.NONE);
    }
}
