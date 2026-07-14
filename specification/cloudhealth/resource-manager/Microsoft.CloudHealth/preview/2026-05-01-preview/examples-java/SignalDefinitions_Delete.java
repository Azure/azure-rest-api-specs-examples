
/**
 * Samples for SignalDefinitions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/SignalDefinitions_Delete.json
     */
    /**
     * Sample code: SignalDefinitions_Delete.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void signalDefinitionsDelete(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.signalDefinitions().delete("online-store-rg", "online-store", "sql-cpu-percent",
            com.azure.core.util.Context.NONE);
    }
}
