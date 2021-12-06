Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-avs_1.0.0-beta.3/sdk/avs/azure-resourcemanager-avs/README.md) on how to add the SDK to your project and authenticate.

```java
/** Samples for CloudLinks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/CloudLinks_CreateOrUpdate.json
     */
    /**
     * Sample code: CloudLinks_CreateOrUpdate.
     *
     * @param manager Entry point to AvsManager.
     */
    public static void cloudLinksCreateOrUpdate(com.azure.resourcemanager.avs.AvsManager manager) {
        manager
            .cloudLinks()
            .define("cloudLink1")
            .withExistingPrivateCloud("group1", "cloud1")
            .withLinkedCloud(
                "/subscriptions/12341234-1234-1234-1234-123412341234/resourceGroups/mygroup/providers/Microsoft.AVS/privateClouds/cloud2")
            .create();
    }
}
```
