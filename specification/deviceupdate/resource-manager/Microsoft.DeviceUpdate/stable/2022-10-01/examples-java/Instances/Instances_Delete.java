import com.azure.core.util.Context;

/** Samples for Instances Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2022-10-01/examples/Instances/Instances_Delete.json
     */
    /**
     * Sample code: Deletes instance.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void deletesInstance(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.instances().delete("test-rg", "contoso", "blue", Context.NONE);
    }
}
