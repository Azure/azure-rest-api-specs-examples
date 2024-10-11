
/**
 * Samples for VirtualMachines Redeploy.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/VirtualMachines/
     * redeployVirtualMachine.json
     */
    /**
     * Sample code: redeployVirtualMachine.
     * 
     * @param manager Entry point to LabServicesManager.
     */
    public static void redeployVirtualMachine(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.virtualMachines().redeploy("testrg123", "testlab", "template", com.azure.core.util.Context.NONE);
    }
}
