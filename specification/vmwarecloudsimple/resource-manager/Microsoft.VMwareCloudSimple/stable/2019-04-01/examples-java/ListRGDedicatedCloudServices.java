/** Samples for DedicatedCloudServices ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListRGDedicatedCloudServices.json
     */
    /**
     * Sample code: ListRGDedicatedCloudServices.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void listRGDedicatedCloudServices(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager
            .dedicatedCloudServices()
            .listByResourceGroup("myResourceGroup", null, null, null, com.azure.core.util.Context.NONE);
    }
}
