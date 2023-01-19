/** Samples for DedicatedCloudNodes GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetDedicatedCloudNode.json
     */
    /**
     * Sample code: GetDedicatedCloudNode.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void getDedicatedCloudNode(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager
            .dedicatedCloudNodes()
            .getByResourceGroupWithResponse("myResourceGroup", "myNode", com.azure.core.util.Context.NONE);
    }
}
