import com.azure.core.util.Context;

/** Samples for ConfigurationProfileAssignments ListByMachineName. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listConfigurationProfileAssignmentsByMachineName.json
     */
    /**
     * Sample code: List configuration profile assignments by resourceGroup and machine.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void listConfigurationProfileAssignmentsByResourceGroupAndMachine(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .configurationProfileAssignments()
            .listByMachineName("myResourceGroupName", "myMachineName", Context.NONE);
    }
}
