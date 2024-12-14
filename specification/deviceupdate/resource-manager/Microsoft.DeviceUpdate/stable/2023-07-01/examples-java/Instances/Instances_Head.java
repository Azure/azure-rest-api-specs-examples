
/**
 * Samples for Instances Head.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/Instances/
     * Instances_Head.json
     */
    /**
     * Sample code: Checks whether instance exists.
     * 
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void checksWhetherInstanceExists(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.instances().headWithResponse("test-rg", "contoso", "blue", com.azure.core.util.Context.NONE);
    }
}
