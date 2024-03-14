
/**
 * Samples for AdvancedThreatProtection Create.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/security/resource-manager/Microsoft.Security/stable/2019-01-01/examples/AdvancedThreatProtection/
     * PutAdvancedThreatProtectionSettings_example.json
     */
    /**
     * Sample code: Creates or updates the Advanced Threat Protection settings on a specified resource.
     * 
     * @param manager Entry point to SecurityManager.
     */
    public static void createsOrUpdatesTheAdvancedThreatProtectionSettingsOnASpecifiedResource(
        com.azure.resourcemanager.security.SecurityManager manager) {
        manager.advancedThreatProtections().define().withExistingResourceId(
            "subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/resourceGroups/SampleRG/providers/Microsoft.Storage/storageAccounts/samplestorageaccount")
            .withIsEnabled(true).create();
    }
}
