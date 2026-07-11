
/**
 * Samples for NetworkSecurityPerimeter GetConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-06-01/NspConfigurationGet.json
     */
    /**
     * Sample code: GetNspConfiguration.
     * 
     * @param manager Entry point to BatchManager.
     */
    public static void getNspConfiguration(com.azure.resourcemanager.batch.BatchManager manager) {
        manager.networkSecurityPerimeters().getConfigurationWithResponse("default-azurebatch-japaneast", "sampleacct",
            "00000000-0000-0000-0000-000000000000.sampleassociation", com.azure.core.util.Context.NONE);
    }
}
