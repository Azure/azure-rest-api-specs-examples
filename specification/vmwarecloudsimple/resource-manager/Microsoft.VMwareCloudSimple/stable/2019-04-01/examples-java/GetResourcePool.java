/** Samples for ResourcePools Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/GetResourcePool.json
     */
    /**
     * Sample code: GetResourcePool.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void getResourcePool(com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager
            .resourcePools()
            .getWithResponse("westus2", "myPrivateCloud", "resgroup-26", com.azure.core.util.Context.NONE);
    }
}
