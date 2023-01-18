/** Samples for IntegrationAccountSchemas List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountSchemas_List.json
     */
    /**
     * Sample code: Get schemas by integration account name.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getSchemasByIntegrationAccountName(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountSchemas()
            .list("testResourceGroup", "<integrationAccountName>", null, null, com.azure.core.util.Context.NONE);
    }
}
