/** Samples for IntegrationServiceEnvironments GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_Get.json
     */
    /**
     * Sample code: Get integration service environment by name.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getIntegrationServiceEnvironmentByName(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationServiceEnvironments()
            .getByResourceGroupWithResponse(
                "testResourceGroup", "testIntegrationServiceEnvironment", com.azure.core.util.Context.NONE);
    }
}
