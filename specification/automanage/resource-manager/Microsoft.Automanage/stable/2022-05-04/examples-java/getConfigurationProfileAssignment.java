import com.azure.core.util.Context;

/** Samples for ConfigurationProfileAssignments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getConfigurationProfileAssignment.json
     */
    /**
     * Sample code: Get a configuration profile assignment.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void getAConfigurationProfileAssignment(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .configurationProfileAssignments()
            .getWithResponse("myResourceGroupName", "default", "myVMName", Context.NONE);
    }
}
