/** Samples for IntegrationAccountBatchConfigurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountBatchConfigurations_Get.json
     */
    /**
     * Sample code: Get a batch configuration.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getABatchConfiguration(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountBatchConfigurations()
            .getWithResponse(
                "testResourceGroup",
                "testIntegrationAccount",
                "testBatchConfiguration",
                com.azure.core.util.Context.NONE);
    }
}
