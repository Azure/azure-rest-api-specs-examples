/** Samples for SkusAvailability List. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/ListAvailabilities.json
     */
    /**
     * Sample code: ListAvailabilities.
     *
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void listAvailabilities(
        com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager.skusAvailabilities().list("westus2", null, com.azure.core.util.Context.NONE);
    }
}
