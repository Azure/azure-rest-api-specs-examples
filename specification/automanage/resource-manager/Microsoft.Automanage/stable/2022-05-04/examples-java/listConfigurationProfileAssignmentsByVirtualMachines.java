import com.azure.core.util.Context;

/** Samples for ConfigurationProfileAssignments ListByVirtualMachines. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listConfigurationProfileAssignmentsByVirtualMachines.json
     */
    /**
     * Sample code: List configuration profile assignments by resourceGroup and virtual machine.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void listConfigurationProfileAssignmentsByResourceGroupAndVirtualMachine(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .configurationProfileAssignments()
            .listByVirtualMachines("myResourceGroupName", "myVMName", Context.NONE);
    }
}
