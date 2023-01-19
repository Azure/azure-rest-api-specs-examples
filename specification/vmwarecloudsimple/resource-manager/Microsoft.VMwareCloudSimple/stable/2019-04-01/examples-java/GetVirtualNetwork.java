/** Samples for VirtualNetworks Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetVirtualNetwork.json
     */
    /**
     * Sample code: GetVirtualNetwork.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void getVirtualNetwork(com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager
            .virtualNetworks()
            .getWithResponse("westus2", "myPrivateCloud", "dvportgroup-19", com.azure.core.util.Context.NONE);
    }
}
