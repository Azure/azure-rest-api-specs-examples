/** Samples for IntegrationAccountAssemblies ListContentCallbackUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountAssemblies_ListContentCallbackUrl.json
     */
    /**
     * Sample code: Get the callback url for an integration account assembly.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getTheCallbackUrlForAnIntegrationAccountAssembly(
        com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountAssemblies()
            .listContentCallbackUrlWithResponse(
                "testResourceGroup", "testIntegrationAccount", "testAssembly", com.azure.core.util.Context.NONE);
    }
}
