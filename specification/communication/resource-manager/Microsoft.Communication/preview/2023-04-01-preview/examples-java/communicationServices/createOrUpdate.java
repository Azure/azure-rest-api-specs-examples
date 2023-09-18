/** Samples for CommunicationServices CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/communication/resource-manager/Microsoft.Communication/preview/2023-04-01-preview/examples/communicationServices/createOrUpdate.json
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
            .withRegion("Global")
            .withExistingResourceGroup("MyResourceGroup")
            .withDataLocation("United States")
            .create();
    }
}
