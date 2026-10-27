
/**
 * Samples for Flows GeneratePassphrase.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/generateFlowPassphrase.json
     */
    /**
     * Sample code: Generate a new compliant streaming passphrase.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void
        generateANewCompliantStreamingPassphrase(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.flows().generatePassphrase("testRG", "testConnection", "testFlow", com.azure.core.util.Context.NONE);
    }
}
