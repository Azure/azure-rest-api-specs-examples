
/**
 * Samples for Authorizations CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/Authorizations_CreateOrUpdate.json
     */
    /**
     * Sample code: Authorizations_CreateOrUpdate.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void authorizationsCreateOrUpdate(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.authorizations().define("authorization1").withExistingPrivateCloud("group1", "cloud1").create();
    }
}
