
import com.azure.resourcemanager.chaos.models.PrivateAccessProperties;
import com.azure.resourcemanager.chaos.models.PublicNetworkAccessOption;

/**
 * Samples for PrivateAccesses CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/
     * PrivateAccesses_CreateOrUpdate_Create_Or_Update_A_Private_Access_Resource_With_Public_Network_Access.json
     */
    /**
     * Sample code: Create or Update a private access resource with publicNetworkAccess.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void createOrUpdateAPrivateAccessResourceWithPublicNetworkAccess(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.privateAccesses().define("myPrivateAccess").withRegion("centraluseuap")
            .withExistingResourceGroup("myResourceGroup")
            .withProperties(new PrivateAccessProperties().withPublicNetworkAccess(PublicNetworkAccessOption.ENABLED))
            .create();
    }
}
