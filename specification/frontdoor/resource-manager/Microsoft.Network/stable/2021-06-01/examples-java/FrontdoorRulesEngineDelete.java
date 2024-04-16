
/**
 * Samples for RulesEngines Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/frontdoor/resource-manager/Microsoft.Network/stable/2021-06-01/examples/FrontdoorRulesEngineDelete.
     * json
     */
    /**
     * Sample code: Delete Rules Engine Configuration.
     * 
     * @param manager Entry point to FrontDoorManager.
     */
    public static void deleteRulesEngineConfiguration(com.azure.resourcemanager.frontdoor.FrontDoorManager manager) {
        manager.rulesEngines().delete("rg1", "frontDoor1", "rulesEngine1", com.azure.core.util.Context.NONE);
    }
}
