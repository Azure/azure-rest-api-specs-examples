/** Samples for Agents CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2023-03-01/examples/Agents_CreateOrUpdate.json
     */
    /**
     * Sample code: Agents_CreateOrUpdate.
     *
     * @param manager Entry point to StorageMoverManager.
     */
    public static void agentsCreateOrUpdate(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager
            .agents()
            .define("examples-agentName")
            .withExistingStorageMover("examples-rg", "examples-storageMoverName")
            .withArcResourceId(
                "/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName")
            .withArcVmUuid("3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9")
            .withDescription("Example Agent Description")
            .create();
    }
}
