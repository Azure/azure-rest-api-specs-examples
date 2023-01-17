/** Samples for IntegrationAccountAgreements Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountAgreements_Delete.json
     */
    /**
     * Sample code: Delete an agreement.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void deleteAnAgreement(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountAgreements()
            .deleteWithResponse(
                "testResourceGroup", "testIntegrationAccount", "testAgreement", com.azure.core.util.Context.NONE);
    }
}
