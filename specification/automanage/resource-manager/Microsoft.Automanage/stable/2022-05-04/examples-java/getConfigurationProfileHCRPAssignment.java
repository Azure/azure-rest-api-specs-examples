import com.azure.core.util.Context;

/** Samples for ConfigurationProfileHcrpAssignments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/getConfigurationProfileHCRPAssignment.json
     */
    /**
     * Sample code: Get a HCRP configuration profile assignment.
     *
     * @param manager Entry point to AutomanageManager.
     */
    public static void getAHCRPConfigurationProfileAssignment(
        com.azure.resourcemanager.automanage.AutomanageManager manager) {
        manager
            .configurationProfileHcrpAssignments()
            .getWithResponse("myResourceGroupName", "myMachineName", "default", Context.NONE);
    }
}
