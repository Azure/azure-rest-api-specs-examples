import com.azure.resourcemanager.maintenance.fluent.models.ConfigurationAssignmentInner;

/** Samples for ConfigurationAssignments CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/maintenance/resource-manager/Microsoft.Maintenance/stable/2023-04-01/examples/ConfigurationAssignments_CreateOrUpdate.json
     */
    /**
     * Sample code: ConfigurationAssignments_CreateOrUpdate.
     *
     * @param manager Entry point to MaintenanceManager.
     */
    public static void configurationAssignmentsCreateOrUpdate(
        com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager
            .configurationAssignments()
            .createOrUpdateWithResponse(
                "examplerg",
                "Microsoft.Compute",
                "virtualMachineScaleSets",
                "smdtest1",
                "workervmConfiguration",
                new ConfigurationAssignmentInner()
                    .withMaintenanceConfigurationId(
                        "/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/configuration1"),
                com.azure.core.util.Context.NONE);
    }
}
