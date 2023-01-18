/** Samples for Capacities Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/deleteCapacity.json
     */
    /**
     * Sample code: Get details of a capacity.
     *
     * @param manager Entry point to PowerBIDedicatedManager.
     */
    public static void getDetailsOfACapacity(
        com.azure.resourcemanager.powerbidedicated.PowerBIDedicatedManager manager) {
        manager.capacities().delete("TestRG", "azsdktest", com.azure.core.util.Context.NONE);
    }
}
