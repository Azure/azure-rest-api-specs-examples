
/**
 * Samples for UserAssignedIdentities ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/IdentityListByResourceGroup.json
     */
    /**
     * Sample code: IdentityListByResourceGroup.
     * 
     * @param manager Entry point to MsiManager.
     */
    public static void identityListByResourceGroup(com.azure.resourcemanager.msi.MsiManager manager) {
        manager.serviceClient().getUserAssignedIdentities().listByResourceGroup("rgName",
            com.azure.core.util.Context.NONE);
    }
}
