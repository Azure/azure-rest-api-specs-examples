
/**
 * Samples for Cluster Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-18-preview/Cluster_Delete_MinimumSet_Gen.json
     */
    /**
     * Sample code: Cluster_Delete_MinimumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void clusterDeleteMinimumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.clusters().delete("rgconfluent", "tvbhdezawspzzfprrnjoxfwtwlp", "mtmberahkmffekuuz",
            "nyfmkuwyeqhkgwehdjakbjheujj", com.azure.core.util.Context.NONE);
    }
}
