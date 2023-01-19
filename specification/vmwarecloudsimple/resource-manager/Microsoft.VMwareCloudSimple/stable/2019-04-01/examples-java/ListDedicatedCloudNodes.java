/** Samples for DedicatedCloudNodes List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListDedicatedCloudNodes.json
     */
    /**
     * Sample code: ListDedicatedCloudNodes.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void listDedicatedCloudNodes(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager.dedicatedCloudNodes().list(null, null, null, com.azure.core.util.Context.NONE);
    }
}
