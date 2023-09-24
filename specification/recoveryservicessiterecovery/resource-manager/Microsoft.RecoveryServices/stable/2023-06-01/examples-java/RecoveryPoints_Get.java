/** Samples for RecoveryPoints Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/recoveryservicessiterecovery/resource-manager/Microsoft.RecoveryServices/stable/2023-06-01/examples/RecoveryPoints_Get.json
     */
    /**
     * Sample code: Gets a recovery point.
     *
     * @param manager Entry point to SiteRecoveryManager.
     */
    public static void getsARecoveryPoint(
        com.azure.resourcemanager.recoveryservicessiterecovery.SiteRecoveryManager manager) {
        manager
            .recoveryPoints()
            .getWithResponse(
                "vault1",
                "resourceGroupPS1",
                "cloud1",
                "cloud_6d224fc6-f326-5d35-96de-fbf51efb3179",
                "f8491e4f-817a-40dd-a90c-af773978c75b",
                "b22134ea-620c-474b-9fa5-3c1cb47708e3",
                com.azure.core.util.Context.NONE);
    }
}
