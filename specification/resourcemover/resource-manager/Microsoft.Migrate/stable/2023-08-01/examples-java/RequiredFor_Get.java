/** Samples for MoveCollections ListRequiredFor. */
public final class Main {
    /*
     * x-ms-original-file: specification/resourcemover/resource-manager/Microsoft.Migrate/stable/2023-08-01/examples/RequiredFor_Get.json
     */
    /**
     * Sample code: RequiredFor_Get.
     *
     * @param manager Entry point to ResourceMoverManager.
     */
    public static void requiredForGet(com.azure.resourcemanager.resourcemover.ResourceMoverManager manager) {
        manager
            .moveCollections()
            .listRequiredForWithResponse(
                "rg1",
                "movecollection1",
                "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/nic1",
                com.azure.core.util.Context.NONE);
    }
}
