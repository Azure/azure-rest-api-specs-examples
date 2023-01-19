/** Samples for Capacities ListSkusForCapacity. */
public final class Main {
    /*
     * x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/listSKUsForExisting.json
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
