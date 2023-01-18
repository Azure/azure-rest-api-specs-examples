/** Samples for IntegrationServiceEnvironments ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_ListByResourceGroup.json
     */
    /**
     * Sample code: List integration service environments by resource group name.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void listIntegrationServiceEnvironmentsByResourceGroupName(
        com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationServiceEnvironments()
            .listByResourceGroup("testResourceGroup", null, com.azure.core.util.Context.NONE);
    }
}
