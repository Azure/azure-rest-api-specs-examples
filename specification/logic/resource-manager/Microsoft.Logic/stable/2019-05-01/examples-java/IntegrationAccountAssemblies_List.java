/** Samples for IntegrationAccountAssemblies List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountAssemblies_List.json
     */
    /**
     * Sample code: List integration account assemblies.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void listIntegrationAccountAssemblies(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountAssemblies()
            .list("testResourceGroup", "testIntegrationAccount", com.azure.core.util.Context.NONE);
    }
}
