import com.azure.core.util.Context;

/** Samples for Instances Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/preview/2022-04-01-preview/examples/Instances/Instances_Get.json
     */
    /**
     * Sample code: Gets list of Instances.
     *
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void getsListOfInstances(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.instances().getWithResponse("test-rg", "contoso", "blue", Context.NONE);
    }
}
