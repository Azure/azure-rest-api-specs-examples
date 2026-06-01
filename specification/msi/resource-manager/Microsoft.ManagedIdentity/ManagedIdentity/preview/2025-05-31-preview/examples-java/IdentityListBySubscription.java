
/**
 * Samples for UserAssignedIdentities List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-31-preview/IdentityListBySubscription.json
     */
    /**
     * Sample code: IdentityListBySubscription.
     * 
     * @param manager Entry point to MsiManager.
     */
    public static void identityListBySubscription(com.azure.resourcemanager.msi.MsiManager manager) {
        manager.serviceClient().getUserAssignedIdentities().list(com.azure.core.util.Context.NONE);
    }
}
