/** Samples for RulesEngines Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorRulesEngineGet.json
     */
    /**
     * Sample code: Get Rules Engine Configuration.
     *
     * @param manager Entry point to FrontDoorManager.
     */
    public static void getRulesEngineConfiguration(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.rulesEngines().getWithResponse("rg1", "frontDoor1", "rulesEngine1", com.azure.core.util.Context.NONE);
    }
}
