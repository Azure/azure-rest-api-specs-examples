/** Samples for DedicatedCloudNodes ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListRGDedicatedCloudNodes.json
     */
    /**
     * Sample code: ListRGDedicatedCloudNodes.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void listRGDedicatedCloudNodes(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager
            .dedicatedCloudNodes()
            .listByResourceGroup("myResourceGroup", null, null, null, com.azure.core.util.Context.NONE);
    }
}
