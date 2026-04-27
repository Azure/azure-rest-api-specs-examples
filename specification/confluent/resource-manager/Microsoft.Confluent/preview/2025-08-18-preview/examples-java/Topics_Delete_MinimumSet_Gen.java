
/**
 * Samples for Topics Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Topics_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: Topics_Delete_MinimumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void topicsDeleteMinimumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.topics().delete("rgconfluent", "dmkqbkbzegenjirw", "flqluwoymahhtfjmx", "xrqfldtrcxvbxxqwbbouosmvnckut",
            "uflu", com.azure.core.util.Context.NONE);
    }
}
