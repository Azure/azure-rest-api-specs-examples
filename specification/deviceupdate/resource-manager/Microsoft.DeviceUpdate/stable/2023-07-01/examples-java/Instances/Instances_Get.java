
/**
 * Samples for Instances Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/deviceupdate/resource-manager/Microsoft.DeviceUpdate/stable/2023-07-01/examples/Instances/
     * Instances_Get.json
     */
    /**
     * Sample code: Gets list of Instances.
     * 
     * @param manager Entry point to DeviceUpdateManager.
     */
    public static void getsListOfInstances(com.azure.resourcemanager.deviceupdate.DeviceUpdateManager manager) {
        manager.instances().getWithResponse("test-rg", "contoso", "blue", com.azure.core.util.Context.NONE);
    }
}
