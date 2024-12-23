
/**
 * Samples for Agents CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/storagemover/resource-manager/Microsoft.StorageMover/stable/2024-07-01/examples/
     * Agents_CreateOrUpdate_MinimumSet.json
     */
    /**
     * Sample code: Agents_CreateOrUpdate_MinimumSet.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void
        agentsCreateOrUpdateMinimumSet(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.agents().define("examples-agentName")
            .withExistingStorageMover("examples-rg", "examples-storageMoverName")
            .withArcResourceId(
                "/subscriptions/60bcfc77-6589-4da2-b7fd-f9ec9322cf95/resourceGroups/examples-rg/providers/Microsoft.HybridCompute/machines/examples-hybridComputeName")
            .withArcVmUuid("3bb2c024-eba9-4d18-9e7a-1d772fcc5fe9").create();
    }
}
