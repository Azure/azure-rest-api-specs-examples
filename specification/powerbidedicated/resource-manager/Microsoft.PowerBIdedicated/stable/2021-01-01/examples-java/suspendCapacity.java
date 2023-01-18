/** Samples for Capacities Suspend. */
public final class Main {
    /*
     * x-ms-original-file: specification/powerbidedicated/resource-manager/Microsoft.PowerBIdedicated/stable/2021-01-01/examples/suspendCapacity.json
     */
    /**
     * Sample code: Suspend capacity.
     *
     * @param manager Entry point to PowerBIDedicatedManager.
     */
    public static void suspendCapacity(com.azure.resourcemanager.powerbidedicated.PowerBIDedicatedManager manager) {
        manager.capacities().suspend("TestRG", "azsdktest", com.azure.core.util.Context.NONE);
    }
}
