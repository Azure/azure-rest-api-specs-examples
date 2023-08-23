/** Samples for CloudLinks CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/CloudLinks_CreateOrUpdate.json
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
