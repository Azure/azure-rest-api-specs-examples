
/**
 * Samples for PrivateLinkResourcesOperation Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01/GetPrivateLinkResources.json
     */
    /**
     * Sample code: Get PrivateLinkResource.
     * 
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void
        getPrivateLinkResource(com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.privateLinkResourcesOperations().getWithResponse("petesting", "pemsi-ecy-rsv2", "backupResource",
            com.azure.core.util.Context.NONE);
    }
}
