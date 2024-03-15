
/**
 * Samples for SoftwareInventories List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2021-05-01-preview/examples/
     * SoftwareInventories/ListBySubscriptionSoftwareInventories_example.json
     */
    /**
     * Sample code: Gets the software inventory of all virtual machines in the subscriptions.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getsTheSoftwareInventoryOfAllVirtualMachinesInTheSubscriptions(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.softwareInventories().list(com.azure.core.util.Context.NONE);
    }
}
