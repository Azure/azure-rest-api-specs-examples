/** Samples for IntegrationAccountPartners List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountPartners_List.json
     */
    /**
     * Sample code: Get partners by integration account name.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getPartnersByIntegrationAccountName(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountPartners()
            .list("testResourceGroup", "testIntegrationAccount", null, null, com.azure.core.util.Context.NONE);
    }
}
