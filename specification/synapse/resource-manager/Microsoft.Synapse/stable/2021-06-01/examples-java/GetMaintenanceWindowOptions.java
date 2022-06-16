import com.azure.core.util.Context;

/** Samples for SqlPoolMaintenanceWindowOptions Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetMaintenanceWindowOptions.json
     */
    /**
     * Sample code: Get list of transparent data encryption configurations of a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getListOfTransparentDataEncryptionConfigurationsOfASQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolMaintenanceWindowOptions()
            .getWithResponse("samplerg", "testworkspace", "testsp", "current", Context.NONE);
    }
}
