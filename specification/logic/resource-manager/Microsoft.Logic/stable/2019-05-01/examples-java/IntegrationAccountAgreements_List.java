/** Samples for IntegrationAccountAgreements List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountAgreements_List.json
     */
    /**
     * Sample code: Get agreements by integration account name.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getAgreementsByIntegrationAccountName(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountAgreements()
            .list("testResourceGroup", "testIntegrationAccount", null, null, com.azure.core.util.Context.NONE);
    }
}
