```java
import com.azure.resourcemanager.avs.models.Sku;

/** Samples for Clusters CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/Clusters_CreateOrUpdate.json
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
            .create();
    }
}
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.
