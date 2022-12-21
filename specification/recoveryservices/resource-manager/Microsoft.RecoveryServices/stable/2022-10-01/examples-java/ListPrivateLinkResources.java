import com.azure.core.util.Context;

/** Samples for PrivateLinkResourcesOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservices/resource-manager/Microsoft.RecoveryServices/stable/2022-10-01/examples/ListPrivateLinkResources.json
     */
    /**
     * Sample code: List PrivateLinkResources.
     *
     * @param manager Entry point to RecoveryServicesManager.
     */
    public static void listPrivateLinkResources(
        com.azure.resourcemanager.recoveryservices.RecoveryServicesManager manager) {
        manager.privateLinkResourcesOperations().list("petesting", "pemsi-ecy-rsv2", Context.NONE);
    }
}
