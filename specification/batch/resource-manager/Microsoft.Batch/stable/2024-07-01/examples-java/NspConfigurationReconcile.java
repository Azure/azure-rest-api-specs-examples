
/**
 * Samples for NetworkSecurityPerimeter ReconcileConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/NspConfigurationReconcile.json
     */
    /**
     * Sample code: ReconcileNspConfiguration.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void reconcileNspConfiguration(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.networkSecurityPerimeters().reconcileConfiguration("default-azurebatch-japaneast", "sampleacct",
            "00000000-0000-0000-0000-000000000000.sampleassociation", com.azure.core.util.Context.NONE);
    }
}
