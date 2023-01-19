/** Samples for DedicatedCloudServices Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/DeleteDedicatedCloudService.json
     */
    /**
     * Sample code: DeleteDedicatedCloudService.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void deleteDedicatedCloudService(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager.dedicatedCloudServices().delete("myResourceGroup", "myService", com.azure.core.util.Context.NONE);
    }
}
