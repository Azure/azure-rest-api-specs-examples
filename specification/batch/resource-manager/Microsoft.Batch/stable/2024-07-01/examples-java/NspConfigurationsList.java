
/**
 * Samples for NetworkSecurityPerimeter ListConfigurations.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/batch/resource-manager/Microsoft.Batch/stable/2024-07-01/examples/NspConfigurationsList.json
     */
    /**
     * Sample code: ListNspConfigurations.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void listNspConfigurations(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.networkSecurityPerimeters().listConfigurations("default-azurebatch-japaneast", "sampleacct",
            com.azure.core.util.Context.NONE);
    }
}
