
/**
 * Samples for Troubleshooters Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/help/resource-manager/Microsoft.Help/preview/2024-03-01-preview/examples/Troubleshooter_Get.json
     */
    /**
     * Sample code: Troubleshooters_Get.
     * 
     * @param manager Entry point to SelfHelpManager.
     */
    public static void troubleshootersGet(com.azure.resourcemanager.selfhelp.SelfHelpManager manager) {
        manager.troubleshooters().getWithResponse(
            "subscriptions/mySubscription/resourcegroups/myresourceGroup/providers/Microsoft.KeyVault/vaults/test-keyvault-rp",
            "abf168ed-1b54-454a-86f6-e4b62253d3b1", com.azure.core.util.Context.NONE);
    }
}
