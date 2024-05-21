
/**
 * Samples for ConfigurationAssignments Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/
     * ConfigurationAssignments_Get.json
     */
    /**
     * Sample code: ConfigurationAssignments_Get.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void configurationAssignmentsGet(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.configurationAssignments().getWithResponse("examplerg", "Microsoft.Compute", "virtualMachineScaleSets",
            "smdtest1", "workervmConfiguration", com.azure.core.util.Context.NONE);
    }
}
