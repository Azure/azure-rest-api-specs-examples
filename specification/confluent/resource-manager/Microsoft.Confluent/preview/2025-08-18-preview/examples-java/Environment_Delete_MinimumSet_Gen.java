
/**
 * Samples for Environment Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Environment_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: Environment_Delete_MinimumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void environmentDeleteMinimumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.environments().delete("rgconfluent", "yetpbmqrfbsanzjzkzdodlcygpj", "quuhiyvpfajfxrqcyxsb",
            com.azure.core.util.Context.NONE);
    }
}
