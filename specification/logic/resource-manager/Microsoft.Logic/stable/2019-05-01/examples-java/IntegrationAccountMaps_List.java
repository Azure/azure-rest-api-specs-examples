/** Samples for IntegrationAccountMaps List. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountMaps_List.json
     */
    /**
     * Sample code: Get maps by integration account name.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getMapsByIntegrationAccountName(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountMaps()
            .list("testResourceGroup", "testIntegrationAccount", null, null, com.azure.core.util.Context.NONE);
    }
}
