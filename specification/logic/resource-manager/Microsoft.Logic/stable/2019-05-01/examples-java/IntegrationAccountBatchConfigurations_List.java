/** Samples for IntegrationAccountBatchConfigurations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountBatchConfigurations_List.json
     */
    /**
     * Sample code: List batch configurations.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void listBatchConfigurations(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountBatchConfigurations()
            .list("testResourceGroup", "testIntegrationAccount", com.azure.core.util.Context.NONE);
    }
}
