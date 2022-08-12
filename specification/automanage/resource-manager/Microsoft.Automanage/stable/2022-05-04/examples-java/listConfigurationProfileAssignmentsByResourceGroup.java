import com.azure.core.util.Context;

/** Samples for ConfigurationProfileAssignments ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listConfigurationProfileAssignmentsByResourceGroup.json
     */
    /**
     * Sample code: List configuration profile assignments by resourceGroup.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void listConfigurationProfileAssignmentsByResourceGroup(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager.configurationProfileAssignments().listByResourceGroup("myResourceGroupName", Context.NONE);
    }
}
