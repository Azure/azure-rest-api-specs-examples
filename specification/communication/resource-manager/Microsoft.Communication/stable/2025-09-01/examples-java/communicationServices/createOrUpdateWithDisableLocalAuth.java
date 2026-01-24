
/**
 * Samples for CommunicationServices CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/
     * communicationServices/createOrUpdateWithDisableLocalAuth.json
     */
    /**
     * Sample code: Create or update resource with DisableLocalAuth.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void createOrUpdateResourceWithDisableLocalAuth(
        com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.communicationServices().define("MyCommunicationResource").withRegion("Global")
            .withExistingResourceGroup("MyResourceGroup").withDataLocation("United States").withDisableLocalAuth(true)
            .create();
    }
}
