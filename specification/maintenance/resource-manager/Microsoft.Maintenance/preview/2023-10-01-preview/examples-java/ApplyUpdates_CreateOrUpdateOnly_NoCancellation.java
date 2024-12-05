
import com.azure.resourcemanager.maintenance.fluent.models.ApplyUpdateInner;

/**
 * Samples for ApplyUpdates CreateOrUpdateOrCancel.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/maintenance/resource-manager/Microsoft.Maintenance/preview/2023-10-01-preview/examples/
     * ApplyUpdates_CreateOrUpdateOnly_NoCancellation.json
     */
    /**
     * Sample code: ApplyUpdates_CreateOrUpdateOnly_NoCancellation.
     * 
     * @param manager Entry point to MaintenanceManager.
     */
    public static void
        applyUpdatesCreateOrUpdateOnlyNoCancellation(com.azure.resourcemanager.maintenance.MaintenanceManager manager) {
        manager.applyUpdates().createOrUpdateOrCancelWithResponse("examplerg", "Microsoft.Compute",
            "virtualMachineScaleSets", "smdtest1", "20230901121200", new ApplyUpdateInner(),
            com.azure.core.util.Context.NONE);
    }
}
