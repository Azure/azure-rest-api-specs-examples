
/**
 * Samples for Topics Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Topics_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Topics_Delete_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void topicsDeleteMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.topics().delete("rgconfluent", "xxoxo", "ohwjl", "llmaybvui", "xnprfffvbjtsnneofwwlpwuzua",
            com.azure.core.util.Context.NONE);
    }
}
