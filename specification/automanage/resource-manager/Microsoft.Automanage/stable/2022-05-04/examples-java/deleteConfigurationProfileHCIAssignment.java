import com.azure.core.util.Context;

/** Samples for ConfigurationProfileHciAssignments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/deleteConfigurationProfileHCIAssignment.json
     */
    /**
     * Sample code: Delete a HCI configuration profile assignment.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void deleteAHCIConfigurationProfileAssignment(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .configurationProfileHciAssignments()
            .deleteWithResponse("myResourceGroupName", "myClusterName", "default", Context.NONE);
    }
}
