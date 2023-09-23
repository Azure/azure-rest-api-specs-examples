/** Samples for SupportedOperatingSystemsOperation Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/SupportedOperatingSystems_Get.json
     */
    /**
     * Sample code: Gets the data of supported operating systems by SRS.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsTheDataOfSupportedOperatingSystemsBySRS(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .supportedOperatingSystemsOperations()
            .getWithResponse("vault1", "resourceGroupPS1", null, com.azure.core.util.Context.NONE);
    }
}
