/** Samples for Capacities ListSkus. */
public final class Main {
    /*
     * x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/listSKUsForNew.json
     */
    /**
     * Sample code: List eligible SKUs for a new capacity.
     *
     * @param manager Entry point to PowerBIDedicatedManager.
     */
    public static void listEligibleSKUsForANewCapacity(
        com.azure.resourcemanager.powerbidedicated.PowerBIDedicatedManager manager) {
        manager.capacities().listSkusWithResponse(com.azure.core.util.Context.NONE);
    }
}
