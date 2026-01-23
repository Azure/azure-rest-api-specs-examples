
import com.azure.resourcemanager.communication.models.PublicNetworkAccess;

/**
 * Samples for CommunicationServices CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/
     * communicationServices/createOrUpdateWithPublicNetworkAccess.json
     */
    /**
     * Sample code: Create or update resource with PublicNetworkAccess.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void createOrUpdateResourceWithPublicNetworkAccess(
        com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.communicationServices().define("MyCommunicationResource").withRegion("Global")
            .withExistingResourceGroup("MyResourceGroup").withDataLocation("United States")
            .withPublicNetworkAccess(PublicNetworkAccess.ENABLED).create();
    }
}
