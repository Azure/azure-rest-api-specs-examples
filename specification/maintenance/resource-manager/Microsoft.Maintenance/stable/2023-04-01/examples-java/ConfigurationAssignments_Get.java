/** Samples for ConfigurationAssignments Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/ConfigurationAssignments_Get.json
     */
    /**
     * Sample code: ConfigurationAssignments_Get.
     *
     * @param manager Entry point to MaintenanceManager.
     */
    public static void configurationAssignmentsGet(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager
            .configurationAssignments()
            .getWithResponse(
                "examplerg",
                "Microsoft.Compute",
                "virtualMachineScaleSets",
                "smdtest1",
                "workervmConfiguration",
                com.azure.core.util.Context.NONE);
    }
}
