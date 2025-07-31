
/**
 * Samples for PrivateLinkResources Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2024-09-01/PrivateLinkResource_Get.json
     */
    /**
     * Sample code: Get Private Link Resource.
     * 
     * @param manager Entry point to RecoveryServicesDataReplicationManager.
     */
    public static void getPrivateLinkResource(
        com.azure.resourcemanager.recoveryservicesdatareplication.RecoveryServicesDataReplicationManager manager) {
        manager.privateLinkResources().getWithResponse("rgswagger_2024-09-01", "4", "d",
            com.azure.core.util.Context.NONE);
    }
}
