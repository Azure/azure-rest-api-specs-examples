
/**
 * Samples for VirtualMachines Start.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/VirtualMachines/
     * startVirtualMachine.json
     */
    /**
     * Sample code: startVirtualMachine.
     * 
     * @param manager Entry point to LabServicesManager.
     */
    public static void startVirtualMachine(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.virtualMachines().start("testrg123", "testlab", "template", com.azure.core.util.Context.NONE);
    }
}
