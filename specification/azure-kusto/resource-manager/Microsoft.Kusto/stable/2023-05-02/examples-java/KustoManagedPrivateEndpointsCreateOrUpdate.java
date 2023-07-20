/** Samples for ManagedPrivateEndpoints CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-05-02/examples/KustoManagedPrivateEndpointsCreateOrUpdate.json
     */
    /**
     * Sample code: KustoManagedPrivateEndpointsCreateOrUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoManagedPrivateEndpointsCreateOrUpdate(
        com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .managedPrivateEndpoints()
            .define("managedPrivateEndpointTest")
            .withExistingCluster("kustorptest", "kustoCluster")
            .withPrivateLinkResourceId(
                "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Storage/storageAccounts/storageAccountTest")
            .withGroupId("blob")
            .withRequestMessage("Please Approve.")
            .create();
    }
}
