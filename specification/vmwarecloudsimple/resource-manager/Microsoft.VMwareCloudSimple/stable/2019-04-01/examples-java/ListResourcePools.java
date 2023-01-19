/** Samples for ResourcePools List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListResourcePools.json
     */
    /**
     * Sample code: ListResourcePools.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void listResourcePools(com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager.resourcePools().list("westus2", "myPrivateCloud", com.azure.core.util.Context.NONE);
    }
}
