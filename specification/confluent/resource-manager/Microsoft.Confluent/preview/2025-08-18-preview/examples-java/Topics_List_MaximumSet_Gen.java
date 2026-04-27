
/**
 * Samples for Topics List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Topics_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: Topics_List_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void topicsListMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.topics().list("rgconfluent", "zkei", "cvgvhjgrodfwwhxkm", "majpwlefqsjqpfezvkvd", 28,
            "nqtivttbasuwnkum", com.azure.core.util.Context.NONE);
    }
}
