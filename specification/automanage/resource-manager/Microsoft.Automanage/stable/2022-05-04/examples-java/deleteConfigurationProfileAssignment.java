import com.azure.core.util.Context;

/** Samples for ConfigurationProfileAssignments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/deleteConfigurationProfileAssignment.json
     */
    /**
     * Sample code: Delete an configuration profile assignment.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void deleteAnConfigurationProfileAssignment(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .configurationProfileAssignments()
            .deleteWithResponse("myResourceGroupName", "default", "myVMName", Context.NONE);
    }
}
