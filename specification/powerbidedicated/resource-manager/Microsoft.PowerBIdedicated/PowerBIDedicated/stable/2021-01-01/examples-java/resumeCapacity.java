
/**
 * Samples for Capacities Resume.
 */
public final class Main {
    /*
     * x-ms-original-file: 2021-01-01/resumeCapacity.json
     */
    /**
     * Sample code: Get details of a capacity.
     * 
     * @param manager Entry point to PowerBIDedicatedManager.
     */
    public static void
        getDetailsOfACapacity(com.azure.resourcemanager.powerbidedicated.PowerBIDedicatedManager manager) {
        manager.capacities().resume("TestRG", "azsdktest", com.azure.core.util.Context.NONE);
    }
}
