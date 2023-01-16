/** Samples for Controller Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/dnc/resource-manager/Microsoft.DelegatedNetwork/stable/2021-03-15/examples/putController.json
     */
    /**
     * Sample code: Create controller.
     *
     * @param manager Entry point to DelegatedNetworkManager.
     */
    public static void createController(com.azure.resourcemanager.delegatednetwork.DelegatedNetworkManager manager) {
        manager
            .controllers()
            .define("testcontroller")
            .withRegion("West US")
            .withExistingResourceGroup("TestRG")
            .create();
    }
}
