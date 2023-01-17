/** Samples for Controller GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/getController.json
     */
    /**
     * Sample code: Get details of a controller.
     *
     * @param manager Entry point to DelegatedNetworkManager.
     */
    public static void getDetailsOfAController(
        com.azure.resourcemanager.delegatednetwork.DelegatedNetworkManager manager) {
        manager
            .controllers()
            .getByResourceGroupWithResponse("TestRG", "testcontroller", com.azure.core.util.Context.NONE);
    }
}
