
/**
 * Samples for Capacities ListSkusForCapacity.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-01-01/listSKUsForExisting.json
     */
    /**
     * Sample code: List eligible SKUs for an existing capacity.
     * 
     * @param manager Entry point to PowerBIDedicatedManager.
     */
    public static void listEligibleSKUsForAnExistingCapacity(
        com.azure.resourcemanager.powerbidedicated.PowerBIDedicatedManager manager) {
        manager.capacities().listSkusForCapacityWithResponse("TestRG", "azsdktest", com.azure.core.util.Context.NONE);
    }
}
