/** Samples for IntegrationAccountSessions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountSessions_Get.json
     */
    /**
     * Sample code: Get an integration account session.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getAnIntegrationAccountSession(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountSessions()
            .getWithResponse("testrg123", "testia123", "testsession123-ICN", com.azure.core.util.Context.NONE);
    }
}
