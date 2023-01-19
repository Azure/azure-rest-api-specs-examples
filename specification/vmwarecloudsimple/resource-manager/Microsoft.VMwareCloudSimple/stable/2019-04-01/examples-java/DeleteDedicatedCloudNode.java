/** Samples for DedicatedCloudNodes Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/DeleteDedicatedCloudNode.json
     */
    /**
     * Sample code: DeleteDedicatedCloudNode.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void deleteDedicatedCloudNode(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager
            .dedicatedCloudNodes()
            .deleteByResourceGroupWithResponse("myResourceGroup", "myNode", com.azure.core.util.Context.NONE);
    }
}
