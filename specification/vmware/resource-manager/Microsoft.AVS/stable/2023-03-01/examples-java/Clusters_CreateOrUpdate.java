import com.azure.resourcemanager.avs.models.Sku;

/** Samples for Clusters CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/Clusters_CreateOrUpdate.json
     */
    /**
     * Sample code: Clusters_CreateOrUpdate.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void clustersCreateOrUpdate(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .clusters()
            .define("cluster1")
            .withExistingPrivateCloud("group1", "cloud1")
            .withSku(new Sku().withName("AV20"))
            .withClusterSize(3)
            .create();
    }
}
