/** Samples for IntegrationAccountAssemblies Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountAssemblies_Delete.json
     */
    /**
     * Sample code: Delete an integration account assembly.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void deleteAnIntegrationAccountAssembly(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountAssemblies()
            .deleteWithResponse(
                "testResourceGroup", "testIntegrationAccount", "testAssembly", com.azure.core.util.Context.NONE);
    }
}
