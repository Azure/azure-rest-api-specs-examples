/** Samples for IntegrationServiceEnvironmentManagedApis Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationServiceEnvironments_ManagedApis_Get.json
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
            .getWithResponse(
                "testResourceGroup",
                "testIntegrationServiceEnvironment",
                "servicebus",
                com.azure.core.util.Context.NONE);
    }
}
