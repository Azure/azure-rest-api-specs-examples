import com.azure.core.util.Context;

/** Samples for PrivateLinkResourcesOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/GetPrivateLinkResources.json
     */
    /**
     * Sample code: Get PrivateLinkResource.
     *
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void getPrivateLinkResource(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager
            .privateLinkResourcesOperations()
            .getWithResponse("petesting", "pemsi-ecy-rsv2", "backupResource", Context.NONE);
    }
}
