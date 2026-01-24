
import com.azure.resourcemanager.communication.models.ManagedServiceIdentity;
import com.azure.resourcemanager.communication.models.ManagedServiceIdentityType;

/**
 * Samples for CommunicationServices CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/communication/resource-manager/Microsoft.Communication/stable/2025-09-01/examples/
     * communicationServices/createOrUpdateWithSystemAssignedIdentity.json
     */
    /**
     * Sample code: Create or update resource with managed identity.
     * 
     * @param manager Entry point to CommunicationManager.
     */
    public static void createOrUpdateResourceWithManagedIdentity(
        com.azure.resourcemanager.communication.CommunicationManager manager) {
        manager.communicationServices().define("MyCommunicationResource").withRegion("Global")
            .withExistingResourceGroup("MyResourceGroup")
            .withIdentity(new ManagedServiceIdentity().withType(ManagedServiceIdentityType.SYSTEM_ASSIGNED))
            .withDataLocation("United States").create();
    }
}
