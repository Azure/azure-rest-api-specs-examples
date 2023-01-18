/** Samples for IntegrationServiceEnvironmentSkus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_Skus.json
     */
    /**
     * Sample code: List integration service environment skus.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void listIntegrationServiceEnvironmentSkus(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationServiceEnvironmentSkus()
            .list("testResourceGroup", "testIntegrationServiceEnvironment", com.azure.core.util.Context.NONE);
    }
}
