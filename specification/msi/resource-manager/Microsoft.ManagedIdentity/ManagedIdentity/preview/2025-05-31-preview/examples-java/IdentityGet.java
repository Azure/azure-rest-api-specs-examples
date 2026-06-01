
/**
 * Samples for UserAssignedIdentities GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/IdentityGet.json
     */
    /**
     * Sample code: IdentityGet.
     * 
     * @param manager Entry point to MsiManager.
     */
    public static void identityGet(com.azure.resourcemanager.msi.MsiManager manager) {
        manager.serviceClient().getUserAssignedIdentities().getByResourceGroupWithResponse("rgName", "resourceName",
            com.azure.core.util.Context.NONE);
    }
}
