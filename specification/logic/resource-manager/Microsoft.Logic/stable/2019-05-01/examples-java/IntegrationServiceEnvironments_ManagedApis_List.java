/** Samples for IntegrationServiceEnvironmentManagedApis List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_ManagedApis_List.json
     */
    /**
     * Sample code: Gets the integration service environment managed Apis.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getsTheIntegrationServiceEnvironmentManagedApis(
        com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationServiceEnvironmentManagedApis()
            .list("testResourceGroup", "testIntegrationServiceEnvironment", com.azure.core.util.Context.NONE);
    }
}
