/** Samples for IntegrationAccountSchemas Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountSchemas_Delete.json
     */
    /**
     * Sample code: Delete a schema by name.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void deleteASchemaByName(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountSchemas()
            .deleteWithResponse(
                "testResourceGroup", "testIntegrationAccount", "testSchema", com.azure.core.util.Context.NONE);
    }
}
