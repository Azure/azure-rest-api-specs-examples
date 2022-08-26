import com.azure.core.util.Context;

/** Samples for VirtualMachines Reimage. */
public final class Main {
    /*
     * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/stable/2022-08-01/examples/VirtualMachines/reimageVirtualMachine.json
     */
    /**
     * Sample code: reimageVirtualMachine.
     *
     * @param manager Entry point to LabServicesManager.
     */
    public static void reimageVirtualMachine(com.azure.resourcemanager.labservices.LabServicesManager manager) {
        manager.virtualMachines().reimage("testrg123", "testlab", "template", Context.NONE);
    }
}
