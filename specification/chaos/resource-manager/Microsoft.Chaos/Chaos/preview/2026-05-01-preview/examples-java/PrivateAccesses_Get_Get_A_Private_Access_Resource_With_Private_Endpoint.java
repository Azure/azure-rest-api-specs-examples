
/**
 * Samples for PrivateAccesses GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * 2026-05-01-preview/PrivateAccesses_Get_Get_A_Private_Access_Resource_With_Private_Endpoint.json
     */
    /**
     * Sample code: Get a private access resource with private endpoint.
     * 
     * @param manager Entry point to ChaosManager.
     */
    public static void
        getAPrivateAccessResourceWithPrivateEndpoint(com.azure.resourcemanager.chaos.ChaosManager manager) {
        manager.privateAccesses().getByResourceGroupWithResponse("myResourceGroup", "myPrivateAccess",
            com.azure.core.util.Context.NONE);
    }
}
