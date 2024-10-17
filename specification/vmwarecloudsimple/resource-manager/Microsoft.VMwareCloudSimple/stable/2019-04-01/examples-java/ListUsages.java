
/**
 * Samples for Usages List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/vmwarecloudsimple/resource-manager/Microsoft.VMwareCloudSimple/stable/2019-04-01/examples/
     * ListUsages.json
     */
    /**
     * Sample code: ListUsages.
     * 
     * @param manager Entry point to VMwareCloudSimpleManager.
     */
    public static void listUsages(com.azure.resourcemanager.vmwarecloudsimple.VMwareCloudSimpleManager manager) {
        manager.usages().list("westus2", null, com.azure.core.util.Context.NONE);
    }
}
