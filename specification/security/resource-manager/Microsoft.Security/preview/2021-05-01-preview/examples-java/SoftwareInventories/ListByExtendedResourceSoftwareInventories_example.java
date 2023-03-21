/** Samples for SoftwareInventories ListByExtendedResource. */
public final class Main {
    /*
     * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-05-01-preview/examples/SoftwareInventories/ListByExtendedResourceSoftwareInventories_example.json
     */
    /**
     * Sample code: Gets the software inventory of the virtual machine.
     *
     * @param manager Entry point to SecurityManager.
     */
    public static void getsTheSoftwareInventoryOfTheVirtualMachine(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager
            .softwareInventories()
            .listByExtendedResource(
                "EITAN-TESTS", "Microsoft.Compute", "virtualMachines", "Eitan-Test1", com.azure.core.util.Context.NONE);
    }
}
