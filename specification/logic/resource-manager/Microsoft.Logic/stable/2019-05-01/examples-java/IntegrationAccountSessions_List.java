/** Samples for IntegrationAccountSessions List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountSessions_List.json
     */
    /**
     * Sample code: Get a list of integration account sessions.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getAListOfIntegrationAccountSessions(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountSessions()
            .list("testrg123", "testia123", null, null, com.azure.core.util.Context.NONE);
    }
}
