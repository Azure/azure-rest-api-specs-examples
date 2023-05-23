/** Samples for RulesEngines ListByFrontDoor. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorRulesEngineList.json
     */
    /**
     * Sample code: List Rules Engine Configurations in a Front Door.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void listRulesEngineConfigurationsInAFrontDoor(
        com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.rulesEngines().listByFrontDoor("rg1", "frontDoor1", com.azure.core.util.Context.NONE);
    }
}
