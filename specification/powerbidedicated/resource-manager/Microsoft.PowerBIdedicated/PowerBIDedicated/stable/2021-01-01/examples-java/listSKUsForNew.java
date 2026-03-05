
/**
 * Samples for Capacities ListSkus.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-01-01/listSKUsForNew.json
     */
    /**
     * Sample code: List eligible SKUs for a new capacity.
     * 
     * @param manager Entry point to PowerBIDedicatedManager.
     */
    public static void
        listEligibleSKUsForANewCapacity(com.azure.resourcemanager.powerbidedicated.PowerBIDedicatedManager manager) {
        manager.capacities().listSkusWithResponse(com.azure.core.util.Context.NONE);
    }
}
