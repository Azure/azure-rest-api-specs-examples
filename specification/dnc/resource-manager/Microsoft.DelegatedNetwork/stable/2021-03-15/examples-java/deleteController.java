/** Samples for Controller Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/deleteController.json
     */
    /**
     * Sample code: Delete controller.
     *
     * @param manager Entry point to DelegatedNetworkManager.
     */
    public static void deleteController(com.azure.resourcemanager.delegatednetwork.DelegatedNetworkManager manager) {
        manager.controllers().delete("TestRG", "testcontroller", com.azure.core.util.Context.NONE);
    }
}
