
/**
 * Samples for IntegrationAccountSessions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountSessions_Delete
     * .json
     */
    /**
     * Sample code: Delete an integration account session.
     * 
     * @param manager Entry point to LogicManager.
     */
    public static void deleteAnIntegrationAccountSession(com.azure.resourcemanager.logic.LogicManager manager) {
        manager.integrationAccountSessions().deleteWithResponse("testrg123", "testia123", "testsession123-ICN",
            com.azure.core.util.Context.NONE);
    }
}
