/** Samples for IntegrationAccountAssemblies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountAssemblies_Get.json
     */
    /**
     * Sample code: Get an integration account assembly.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getAnIntegrationAccountAssembly(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountAssemblies()
            .getWithResponse(
                "testResourceGroup", "testIntegrationAccount", "testAssembly", com.azure.core.util.Context.NONE);
    }
}
