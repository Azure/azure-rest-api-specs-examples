
/**
 * Samples for SignalDefinitions Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/SignalDefinitions_Get.json
     */
    /**
     * Sample code: SignalDefinitions_Get.
     * 
     * @param manager Entry point to CloudHealthManager.
     */
    public static void signalDefinitionsGet(com.azure.resourcemanager.cloudhealth.CloudHealthManager manager) {
        manager.signalDefinitions().getWithResponse("online-store-rg", "online-store", "sql-cpu-percent",
            com.azure.core.util.Context.NONE);
    }
}
