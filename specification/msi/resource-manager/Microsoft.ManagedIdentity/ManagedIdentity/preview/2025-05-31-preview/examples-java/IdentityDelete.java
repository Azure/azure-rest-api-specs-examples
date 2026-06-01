
/**
 * Samples for UserAssignedIdentities Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/IdentityDelete.json
     */
    /**
     * Sample code: IdentityDelete.
     * 
     * @param manager Entry point to MsiManager.
     */
    public static void identityDelete(com.azure.resourcemanager.msi.MsiManager manager) {
        manager.serviceClient().getUserAssignedIdentities().deleteWithResponse("rgName", "resourceName",
            com.azure.core.util.Context.NONE);
    }
}
