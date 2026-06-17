
import com.azure.resourcemanager.chaos.models.PrivateAccessProperties;

/**
 * Samples for PrivateAccesses CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-05-01-preview/PrivateAccesses_CreateOrUpdate_Create_Or_Update_A_Private_Access_Resource.json
     */
    /**
     * Sample code: Create or Update a private access resource.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void createOrUpdateAPrivateAccessResource(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.privateAccesses().define("myPrivateAccess").withRegion("centraluseuap")
            .withExistingResourceGroup("myResourceGroup").withProperties(new PrivateAccessProperties()).create();
    }
}
