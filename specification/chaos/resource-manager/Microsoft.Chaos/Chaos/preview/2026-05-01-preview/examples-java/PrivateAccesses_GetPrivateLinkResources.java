
/**
 * Samples for PrivateAccesses GetPrivateLinkResources.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/PrivateAccesses_GetPrivateLinkResources.json
     */
    /**
     * Sample code: List all possible private link resources under private access resource.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void listAllPossiblePrivateLinkResourcesUnderPrivateAccessResource(
        com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.privateAccesses().getPrivateLinkResourcesWithResponse("myResourceGroup", "myPrivateAccess",
            com.azure.core.util.Context.NONE);
    }
}
