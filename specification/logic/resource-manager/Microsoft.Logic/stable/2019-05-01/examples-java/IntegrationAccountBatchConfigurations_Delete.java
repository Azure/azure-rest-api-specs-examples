/** Samples for IntegrationAccountBatchConfigurations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountBatchConfigurations_Delete.json
     */
    /**
     * Sample code: Delete a batch configuration.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void deleteABatchConfiguration(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountBatchConfigurations()
            .deleteWithResponse(
                "testResourceGroup",
                "testIntegrationAccount",
                "testBatchConfiguration",
                com.azure.core.util.Context.NONE);
    }
}
