
/**
 * Samples for SoftwareInventories Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/preview/2021-05-01-preview/examples/
     * SoftwareInventories/GetSoftware_example.json
     */
    /**
     * Sample code: Gets a single software data of the virtual machine.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void
        getsASingleSoftwareDataOfTheVirtualMachine(com.azure.resourcemanager.security.SecurityManager manager) {
        manager.softwareInventories().getWithResponse("EITAN-TESTS", "Microsoft.Compute", "virtualMachines",
            "Eitan-Test1", "outlook_16.0.10371.20060", com.azure.core.util.Context.NONE);
    }
}
