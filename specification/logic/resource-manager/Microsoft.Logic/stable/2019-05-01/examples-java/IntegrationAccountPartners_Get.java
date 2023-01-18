/** Samples for IntegrationAccountPartners Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountPartners_Get.json
     */
    /**
     * Sample code: Get partner by name.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getPartnerByName(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountPartners()
            .getWithResponse(
                "testResourceGroup", "testIntegrationAccount", "testPartner", com.azure.core.util.Context.NONE);
    }
}
