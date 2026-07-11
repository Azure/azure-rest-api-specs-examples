
/**
 * Samples for PrivateLinkResourcesOperation List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/ListPrivateLinkResources.json
     */
    /**
     * Sample code: List PrivateLinkResources.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void
        listPrivateLinkResources(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.privateLinkResourcesOperations().list("petesting", "pemsi-ecy-rsv2", com.azure.core.util.Context.NONE);
    }
}
