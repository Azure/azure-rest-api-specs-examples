
/**
 * Samples for Get IntegrationRuntimeStop.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/
     * IntegrationRuntimes_Stop_OperationStatus.json
     */
    /**
     * Sample code: Get integration runtime operation status.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void getIntegrationRuntimeOperationStatus(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.gets().integrationRuntimeStopWithResponse("drage-felles-prod-rg", "felles-prod-synapse-workspace",
            "SSIS-intergrationRuntime-Drage", "5752dcdf918e4aecb941245ddf6ebb83", com.azure.core.util.Context.NONE);
    }
}
