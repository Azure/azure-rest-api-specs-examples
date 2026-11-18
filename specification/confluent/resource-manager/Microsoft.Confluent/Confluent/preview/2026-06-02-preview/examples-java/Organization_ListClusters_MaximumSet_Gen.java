
/**
 * Samples for Organization ListClusters.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-02-preview/Organization_ListClusters_MaximumSet_Gen.json
     */
    /**
     * Sample code: Organization_ListClusters_MaximumSet.
     * 
     * @param manager Entry point to ConfluentManager.
     */
    public static void
        organizationListClustersMaximumSet(com.azure.resourcemanager.confluent.ConfluentManager manager) {
        manager.organizations().listClusters("rgconfluent", "hpinjsodpkprhbvpzh", "qjeouprbl", 24,
            "esiyyipdkqikzcedkyrjnqvsbf", com.azure.core.util.Context.NONE);
    }
}
