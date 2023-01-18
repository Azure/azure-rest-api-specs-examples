/** Samples for IntegrationAccountMaps Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/logic/resource-manager/Microsoft.Logic/stable/2019-05-01/examples/IntegrationAccountMaps_Delete.json
     */
    /**
     * Sample code: Delete a map.
     *
     * @param manager Entry point to LogicManager.
     */
    public static void deleteAMap(com.azure.resourcemanager.logic.LogicManager manager) {
        manager
            .integrationAccountMaps()
            .deleteWithResponse(
                "testResourceGroup", "testIntegrationAccount", "testMap", com.azure.core.util.Context.NONE);
    }
}
