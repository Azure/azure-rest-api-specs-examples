
/**
 * Samples for Cluster Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Cluster_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: Cluster_Delete_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void clusterDeleteMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.clusters().delete("rgconfluent", "rwzpoelzgevhnkrvyqy", "gnijsroqxwwyyariafdnmkc", "zsvnfsirukovzkth",
            com.azure.core.util.Context.NONE);
    }
}
