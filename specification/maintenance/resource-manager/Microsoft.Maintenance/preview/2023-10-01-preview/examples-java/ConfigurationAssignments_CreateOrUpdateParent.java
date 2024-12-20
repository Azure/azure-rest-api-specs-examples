
import com.azure.resourcemanager.maintenance.fluent.models.ConfigurationAssignmentInner;

/**
 * Samples for ConfigurationAssignments CreateOrUpdateParent.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/
     * ConfigurationAssignments_CreateOrUpdateParent.json
     */
    /**
     * Sample code: ConfigurationAssignments_CreateOrUpdateParent.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void
        configurationAssignmentsCreateOrUpdateParent(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.configurationAssignments().createOrUpdateParentWithResponse("examplerg", "Microsoft.Compute",
            "virtualMachineScaleSets", "smdtest1", "virtualMachines", "smdvm1", "workervmPolicy",
            new ConfigurationAssignmentInner().withMaintenanceConfigurationId(
                "/subscriptions/5b4b650e-28b9-4790-b3ab-ddbd88d727c4/resourcegroups/examplerg/providers/Microsoft.Maintenance/maintenanceConfigurations/policy1"),
            com.azure.core.util.Context.NONE);
    }
}
