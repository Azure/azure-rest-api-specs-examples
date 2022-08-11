import com.azure.core.util.Context;

/** Samples for ConfigurationProfileHciAssignments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getConfigurationProfileHCIAssignment.json
     */
    /**
     * Sample code: Get a HCI configuration profile assignment.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void getAHCIConfigurationProfileAssignment(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .configurationProfileHciAssignments()
            .getWithResponse("myResourceGroupName", "myClusterName", "default", Context.NONE);
    }
}
