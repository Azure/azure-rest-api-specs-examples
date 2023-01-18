/** Samples for IntegrationAccountAgreements Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountAgreements_Get.json
     */
    /**
     * Sample code: Get agreement by name.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getAgreementByName(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountAgreements()
            .getWithResponse(
                "testResourceGroup", "testIntegrationAccount", "testAgreement", com.azure.core.util.Context.NONE);
    }
}
