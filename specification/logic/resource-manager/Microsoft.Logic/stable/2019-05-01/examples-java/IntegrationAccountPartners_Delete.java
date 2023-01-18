/** Samples for IntegrationAccountPartners Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountPartners_Delete.json
     */
    /**
     * Sample code: Delete a partner.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void deleteAPartner(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountPartners()
            .deleteWithResponse(
                "testResourceGroup", "testIntegrationAccount", "testPartner", com.azure.core.util.Context.NONE);
    }
}
