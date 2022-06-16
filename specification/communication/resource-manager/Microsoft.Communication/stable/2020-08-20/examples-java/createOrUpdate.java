/** Samples for CommunicationService CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/stable/2020-08-20/examples/createOrUpdate.json
     */
    /**
     * Sample code: Create or update resource.
     *
     * @param manager Entry point to CommunicationManager.
     */
    public static void createOrUpdateResource(com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager
            .communicationServices()
            .define("MyCommunicationResource")
            .withExistingResourceGroup("MyResourceGroup")
            .withRegion("Global")
            .withDataLocation("United States")
            .create();
    }
}
