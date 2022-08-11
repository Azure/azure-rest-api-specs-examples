import com.azure.core.util.Context;

/** Samples for ConfigurationProfileHcrpAssignments Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/deleteConfigurationProfileHCRPAssignment.json
     */
    /**
     * Sample code: Delete a HCRP configuration profile assignment.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void deleteAHCRPConfigurationProfileAssignment(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .configurationProfileHcrpAssignments()
            .deleteWithResponse("myResourceGroupName", "myMachineName", "default", Context.NONE);
    }
}
