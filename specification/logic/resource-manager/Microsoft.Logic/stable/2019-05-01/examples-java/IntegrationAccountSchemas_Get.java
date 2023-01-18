/** Samples for IntegrationAccountSchemas Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountSchemas_Get.json
     */
    /**
     * Sample code: Get schema by name.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getSchemaByName(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountSchemas()
            .getWithResponse(
                "testResourceGroup", "testIntegrationAccount", "testSchema", com.azure.core.util.Context.NONE);
    }
}
