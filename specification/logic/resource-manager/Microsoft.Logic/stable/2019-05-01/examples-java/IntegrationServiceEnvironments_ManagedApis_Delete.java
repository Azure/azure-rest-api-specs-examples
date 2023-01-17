/** Samples for IntegrationServiceEnvironmentManagedApis Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_ManagedApis_Delete.json
     */
    /**
     * Sample code: Deletes the integration service environment managed Apis.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void deletesTheIntegrationServiceEnvironmentManagedApis(
        com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationServiceEnvironmentManagedApis()
            .delete(
                "testResourceGroup",
                "testIntegrationServiceEnvironment",
                "servicebus",
                com.azure.core.util.Context.NONE);
    }
}
