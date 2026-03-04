
import com.azure.resourcemanager.maintenance.fluent.models.ApplyUpdateInner;
import com.azure.resourcemanager.maintenance.models.UpdateStatus;

/**
 * Samples for ApplyUpdates CreateOrUpdateOrCancel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2023-10-01-preview/ApplyUpdates_CreateOrUpdate_CancelMaintenance.json
     */
    /**
     * Sample code: ApplyUpdates_CreateOrUpdateOrCancel.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void
        applyUpdatesCreateOrUpdateOrCancel(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.applyUpdates().createOrUpdateOrCancelWithResponse("examplerg", "Microsoft.Maintenance",
            "maintenanceConfigurations", "maintenanceConfig1", "20230901121200",
            new ApplyUpdateInner().withStatus(UpdateStatus.CANCEL), com.azure.core.util.Context.NONE);
    }
}
