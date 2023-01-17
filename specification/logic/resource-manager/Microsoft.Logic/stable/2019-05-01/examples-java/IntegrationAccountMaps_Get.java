/** Samples for IntegrationAccountMaps Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountMaps_Get.json
     */
    /**
     * Sample code: Get map by name.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void getMapByName(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountMaps()
            .getWithResponse(
                "testResourceGroup", "testIntegrationAccount", "testMap", com.azure.core.util.Context.NONE);
    }
}
