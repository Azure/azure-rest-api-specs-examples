
/**
 * Samples for Topics Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Topics_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: Topics_Get_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void topicsGetMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.topics().getWithResponse("rgconfluent", "mwvtthpz", "gjcsgothfog", "cbgic", "bspwihoyrewjny",
            com.azure.core.util.Context.NONE);
    }
}
