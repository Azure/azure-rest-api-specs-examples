
/**
 * Samples for AdvancedThreatProtection Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2019-01-01/examples/AdvancedThreatProtection/
     * GetAdvancedThreatProtectionSettings_example.json
     */
    /**
     * Sample code: Gets the Advanced Threat Protection settings for the specified resource.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void getsTheAdvancedThreatProtectionSettingsForTheSpecifiedResource(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.advancedThreatProtections().getWithResponse(
            "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Storage/storageAccounts/samplestorageaccount",
            com.azure.core.util.Context.NONE);
    }
}
