import com.azure.core.util.Context;

/** Samples for Instances Head. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/Instances/Instances_Head.json
     */
    /**
     * Sample code: Checks whether instance exists.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void checksWhetherInstanceExists(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.instances().headWithResponse("test-rg", "contoso", "blue", Context.NONE);
    }
}
