/** Samples for PrivateClouds List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListPrivateCloudInLocation.json
     */
    /**
     * Sample code: ListPrivateCloudInLocation.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void listPrivateCloudInLocation(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager.privateClouds().list("eastus", com.azure.core.util.Context.NONE);
    }
}
