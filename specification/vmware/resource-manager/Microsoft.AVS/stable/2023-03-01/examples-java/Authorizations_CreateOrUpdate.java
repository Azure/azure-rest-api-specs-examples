/** Samples for Authorizations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2023-03-01/examples/Authorizations_CreateOrUpdate.json
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
